class Solution {
    fun getNoZeroIntegers(n: Int): IntArray {
        for (i in 1 until n) {
            if (noZero(i) && noZero(n - i)) {
                return intArrayOf(i, n - i)
            }
        }
        return IntArray(0)
    }

    tailrec fun noZero(n: Int): Boolean {
        return when {
            n % 10 == 0 -> false
            n < 10 -> true
            else -> noZero(n / 10)
        }
    }
}