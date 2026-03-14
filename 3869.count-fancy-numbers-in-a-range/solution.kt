class Solution {
    private lateinit var s: String
    private lateinit var goodSums: Set<Int>
    private lateinit var memo: Array<Array<Array<Array<Array<LongArray>>>>>
    
    fun countFancy(l: Long, r: Long): Long {
        return solve(r) - solve(l - 1)        
    }

    private fun solve(n: Long): Long {
        if (n < 1) {
            return 0L
        }
        s = n.toString()
        val maxS = 9 * s.length

        goodSums = (1..maxS).filter { isGood(it) }.toSet()

        memo = Array(s.length) { Array(11) { Array(4) { Array(maxS + 1) { Array(2) { LongArray(2) { -1L } } } } } }

        return dp(0, 10, 3, 0, 1, 0)
    }

    private fun dp(idx: Int, prev: Int, dir: Int, sum: Int, isTight: Int, isStarted: Int): Long {
        if (idx == s.length) {
            return when {
                isStarted == 0 -> 0L
                dir != 2 || sum in goodSums -> 1L
                else -> 0L
            }
        }
        if (memo[idx][prev][dir][sum][isTight][isStarted] != -1L) {
            return memo[idx][prev][dir][sum][isTight][isStarted]
        }

        var ans = 0L
        val limit = if (isTight == 1) s[idx] - '0' else 9

        for (d in 0..limit) {
            val newTight = if (isTight == 1 && d == limit) 1 else 0
            if (isStarted == 0) {
                if (d == 0) {
                    ans += dp(idx + 1, 10, 3, 0, newTight, 0)
                } else {
                    ans += dp(idx + 1, d, 3, d, newTight, 1)
                }
            } else {
                val newDir = getNextDir(dir, prev, d)
                ans += dp(idx + 1, d, newDir, sum + d, newTight, 1)
            }
        }
        memo[idx][prev][dir][sum][isTight][isStarted] = ans
        return ans
    }

    private fun getNextDir(dir: Int, prev: Int, d: Int): Int {
        return when(dir) {
            0 -> if (d > prev) 0 else 2
            1 -> if (d < prev) 1 else 2
            3 -> when {
                d > prev -> 0
                d < prev -> 1
                else -> 2
            }
            else -> 2
        }
    }
    
    private fun isGood(k: Int): Boolean {
        if (k < 10) {
            return true
        }
        
        val digits = k.toString().map { it - '0' }
        var isIncreasing = true
        var isDecreasing = true

        for (i in 0 until digits.lastIndex) {
            if (digits[i] >= digits[i + 1]) {
                isIncreasing = false
            }
            if (digits[i] <= digits[i + 1]) {
                isDecreasing = false
            }
        }
        return isIncreasing || isDecreasing
    }
}