class Solution {
    fun minimumCost(nums: IntArray): Int {
        var min1 = Int.MAX_VALUE
        var min2 = Int.MAX_VALUE
        for (i in 1..nums.lastIndex) {
            val n = nums[i]
            if (n < min1) {
                min2 = min1
                min1 = n
            } else if (n < min2) {
                min2 = n
            }
        }
        return nums.first() + min1 + min2
    }
}