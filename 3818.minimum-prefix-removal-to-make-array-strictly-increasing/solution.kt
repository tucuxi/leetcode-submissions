class Solution {
    fun minimumPrefixLength(nums: IntArray): Int {
        for (i in nums.lastIndex downTo 1) {
            if (nums[i] <= nums[i - 1]) {
                return i
            }
        }
        return 0
    }
}