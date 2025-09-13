class Solution {
    fun makeTheIntegerZero(num1: Int, num2: Int): Int {
        var i = 1
        while (true) {
            val b = num1.toLong() - i * num2.toLong()
            if (b < i) {
                return -1
            }
            if (b.countOneBits() <= i) {
                return i
            }
            i += 1
        }
    }
}