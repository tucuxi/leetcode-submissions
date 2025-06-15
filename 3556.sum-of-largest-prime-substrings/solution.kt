class Solution {
    fun sumOfLargestPrimes(s: String): Long {
        val primes = mutableSetOf<Long>()
        for (i in s.indices) {
            for (j in i + 1 .. s.length) {
                val t = s.substring(i, j)
                val n = t.toLong()
                if (isPrime(n)) {
                    primes.add(n)
                }
            }
        }
        return primes.sortedDescending().take(3).sum()
    }

    fun isPrime(n: Long): Boolean {
        return when(n) {
            1L -> false
            2L -> true
            else -> {
                if (n % 2 == 0L) {
                    return false
                }
                var i = 3L
                while (i * i <= n) {
                    if (n % i == 0L) {
                        return false
                    }
                    i += 2
                }
                true
            }
        }
    }
}