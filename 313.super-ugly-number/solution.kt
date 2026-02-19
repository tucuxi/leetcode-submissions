class Solution {
    fun nthSuperUglyNumber(n: Int, primes: IntArray): Int {
        val res = LongArray(n)
        res[0] = 1

        val h = IntArray(primes.size)

        for (i in 1 until n) {
            res[i] = Long.MAX_VALUE
            for (j in primes.indices) {
                res[i] = minOf(res[i], primes[j].toLong() * res[h[j]])
            }
            for (j in h.indices) {
                if (res[i] == primes[j] * res[h[j]]) {
                    h[j]++
                }
            }
        }

        return res[n - 1].toInt()
    }
}