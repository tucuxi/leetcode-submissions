const val MOD = 1_000_000_007

class Solution {
    fun concatenatedBinary(n: Int): Int {
        var res = 0L
        var l = 0
        for (i in 1..n) {
            if (i and i - 1 == 0) {
                l++
            }
            res = ((res shl l) + i) % MOD
        }
        return res.toInt()
    }
}