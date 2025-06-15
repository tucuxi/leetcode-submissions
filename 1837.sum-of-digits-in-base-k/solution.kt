class Solution {
    fun sumBase(n: Int, k: Int): Int {
        var m = n
        var r = 0
        while (m > 0) {
            r += m % k
            m /= k
        }
        return r
    }
}