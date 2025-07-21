class Solution {
    fun checkPrimeFrequency(nums: IntArray): Boolean {
        val frequencies = nums.asSequence().groupingBy { it }.eachCount()
        return frequencies.any { (_, f) -> isPrime(f) }
    }

    fun isPrime(n: Int): Boolean {
        return when {
            n == 1 -> false
            n == 2 -> true
            n % 2 == 0 -> false
            else -> {
                var f = 3
                while (f * f <= n) {
                    if (n % f == 0) {
                        return false
                    }
                    f += 2
                }
                true
            }
        }
    }
}