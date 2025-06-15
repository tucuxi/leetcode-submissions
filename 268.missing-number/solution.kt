class Solution {
    fun missingNumber(nums: IntArray) =
        nums.size * (nums.size + 1) / 2 - nums.sum()
}