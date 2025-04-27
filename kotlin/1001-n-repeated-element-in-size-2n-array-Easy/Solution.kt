class Solution {
    fun repeatedNTimes(nums: IntArray): Int {
        for (i in 2..nums.lastIndex) {
            if (nums[i] == nums[i - 1] || nums[i] == nums[i - 2]) {
                return nums[i]
            }
        }
        return nums[0]
    }
}