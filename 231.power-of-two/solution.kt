class Solution {
    fun isPowerOfTwo(n: Int): Boolean {
        if (n <= 0) {
            return false
        }
        var p = n
        while (p and 1 == 0) {
            p = p shr 1
        }
        return p == 1
    }
}