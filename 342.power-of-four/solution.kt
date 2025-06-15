class Solution {
    fun isPowerOfFour(n: Int): Boolean {
        var m = n
        while (m > 1) {
            if (m % 4 != 0) return false
            m /= 4
        }
        return m == 1
    }
}