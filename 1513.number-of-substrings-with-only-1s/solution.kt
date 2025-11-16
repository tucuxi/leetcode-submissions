const val MOD = 1_000_000_007

class Solution {
    fun numSub(s: String): Int {
        var res = 0
        var run = 0

        s.forEach { ch ->
            if (ch == '0') {
                run = 0
            } else {
                run++
                res = (res + run) % MOD
            }
        }

        return res
    }
}