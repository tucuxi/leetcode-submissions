const val MOD = 1_000_000_007L

class Solution {
    fun sumOfNumbers(l: Int, r: Int, k: Int): Int {
        val n = (r - l + 1).toLong()
        val sumDigits = (l + r) * n / 2
        val nPow = power(n, (k - 1).toLong(), MOD)
        val tenPow = power(10L, k.toLong(), MOD)
        val geometricSum = (tenPow - 1 + MOD) % MOD
        val inv9 = power(9L, MOD - 2, MOD)
        val repUnit = (geometricSum * inv9) % MOD
        val result = (sumDigits % MOD * nPow % MOD * repUnit % MOD)
        return result.toInt()        
    }

    private fun power(base: Long, exp: Long, mod: Long): Long {
        var res = 1L
        var b = base
        var e = exp
        while (e > 0) {
            if (e % 2 == 1L) res = (res * b) % mod
            b = (b * b) % mod
            e /= 2
        }
        return res
    }

}