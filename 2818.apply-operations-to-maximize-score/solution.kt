const val MOD = 1_000_000_007

class Solution {
    fun maximumScore(nums: List<Int>, k: Int): Int {
        val n = nums.size
        val primeScores = computePrimeScores(nums)
        
        val left = IntArray(n) { -1 }
        val right = IntArray(n) { n }
        val stack = Stack<Int>()
        
        // Find next greater element to the left (strictly greater)
        for (i in 0 until n) {
            while (stack.isNotEmpty() && primeScores[stack.peek()] < primeScores[i]) {
                stack.pop()
            }
            if (stack.isNotEmpty()) {
                left[i] = stack.peek()
            }
            stack.push(i)
        }
        
        stack.clear()
        
        // Find next greater element to the right (greater or equal)
        for (i in n - 1 downTo 0) {
            while (stack.isNotEmpty() && primeScores[stack.peek()] <= primeScores[i]) {
                stack.pop()
            }
            if (stack.isNotEmpty()) {
                right[i] = stack.peek()
            }
            stack.push(i)
        }
        
        val elements = mutableListOf<Pair<Int, Long>>()
        for (i in 0 until n) {
            val l = i - left[i]
            val r = right[i] - i
            val count = l.toLong() * r.toLong()
            elements.add(Pair(nums[i], count))
        }
        
        // Sort elements in descending order of their values
        elements.sortWith(compareByDescending<Pair<Int, Long>> { it.first })
        
        var remainingK = k.toLong()
        var result = 1L
        
        for ((num, cnt) in elements) {
            if (remainingK <= 0) break
            val take = minOf(cnt, remainingK)
            result = (result * modPow(num.toLong(), take, MOD)) % MOD
            remainingK -= take
        }
        
        return result.toInt()
    }
    
    private fun computePrimeScores(nums: List<Int>): IntArray {
        val maxNum = nums.maxOrNull() ?: 0
        val spf = IntArray(maxNum + 1) { it }
        
        // Sieve of Eratosthenes to compute smallest prime factors
        for (i in 2..maxNum) {
            if (spf[i] == i) {
                if (i.toLong() * i > maxNum) continue
                for (j in i * i..maxNum step i) {
                    if (spf[j] == j) {
                        spf[j] = i
                    }
                }
            }
        }
        
        val primeScores = IntArray(nums.size)
        for (i in nums.indices) {
            var x = nums[i]
            var score = 0
            if (x == 1) {
                primeScores[i] = 0
                continue
            }
            var lastPrime = 0
            while (x > 1) {
                val p = spf[x]
                if (p != lastPrime) {
                    score++
                    lastPrime = p
                }
                x /= p
            }
            primeScores[i] = score
        }
        return primeScores
    }
    
    private fun modPow(a: Long, b: Long, mod: Int): Long {
        var result = 1L
        var base = a % mod
        var exponent = b
        while (exponent > 0) {
            if (exponent % 2 == 1L) {
                result = (result * base) % mod
            }
            base = (base * base) % mod
            exponent /= 2
        }
        return result
    }
}