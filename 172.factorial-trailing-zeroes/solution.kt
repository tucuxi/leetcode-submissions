class Solution {
    fun trailingZeroes(n: Int): Int {
        var p = n
        var c = 0
        while (p > 0) {
            p /= 5
            c += p
        }
        return c
    }
}