class Solution {
    fun smallestDivisor(nums: IntArray, threshold: Int): Int {
        var l = 1
        var r = 1000000
        while (l < r) {
            val divisor = l + (r - l) / 2
            val sum = nums.sumOf { (it + divisor - 1) / divisor }
            if (sum <= threshold) {
                r = divisor
            } else {
                l = divisor + 1
            }
        }
        return l
    }
}