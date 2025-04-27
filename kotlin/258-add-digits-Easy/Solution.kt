class Solution {
    fun addDigits(num: Int): Int {
        var p = num
        while (p >= 10) {
            var q = 0
            while (p > 0) {
                q += p % 10
                p /= 10
            }
            p = q
        }
        return p
    }
}