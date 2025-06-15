class Solution {
    fun countPrimes(n: Int): Int {
        if (n < 2) return 0
        val nonPrimes = BooleanArray(n)
        for (i in 2 until n) {
            for (j in 2 * i until n step i) {
                nonPrimes[j] = true
            }
        }
        return n - nonPrimes.count { it } - 2
    }
}