class Solution {
    fun maxProduct(n: Int): Int {
        var max1 = 0
        var max2 = 0
        var n = n
        while (n > 0) {
            val d = n % 10
            if (d > max1) {
                max2 = max1
                max1 = d
            } else if (d > max2) {
                max2 = d
            }
            n /= 10
        }
        return max1 * max2
    }
}