class Solution {
    fun countSubarrays(nums: IntArray): Int {
        return (1 until nums.lastIndex)
            .map { i -> if (2 * (nums[i - 1] + nums[i + 1]) == nums[i]) 1 else 0 }
            .sum()
    }
}