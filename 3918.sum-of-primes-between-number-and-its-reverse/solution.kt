class Solution {
    fun sumOfPrimesInRange(n: Int): Int {
        val r = reverse(n)
        val p = minOf(n, r)
        val q = maxOf(n, r)

        return (p..q).filter { isPrime(it) }.sum()
    }

    fun reverse(n: Int): Int {
        var r = 0
        var i = n

        while (i > 0) {
            r = 10*r + i%10
            i /= 10
        }
        return r
    }

    fun isPrime(n: Int): Boolean {
        return when {
            n == 1 -> false
            n == 2 -> true
            n%2 == 0 -> false
            else -> isPrimeOdd3Plus(n)
        }
    }

    fun isPrimeOdd3Plus(n: Int): Boolean {
        var i = 3

        while (i*i <= n) {
            if (n%i == 0) {
                return false
            }
            i += 2
        }
        return true
    }
}