class Solution {
    fun gcdOfOddEvenSums(n: Int): Int {
        val sum = n * (2 * n + 1)
        val sumOdd = (sum - n) / 2
        val sumEven = sumOdd + n
        return gcd(sumOdd, sumEven)
    }

    tailrec fun gcd(a: Int, b: Int): Int =
        when {
            a == 0 -> b
            b == 0 -> a
            else -> gcd(b, a % b)
        }
}