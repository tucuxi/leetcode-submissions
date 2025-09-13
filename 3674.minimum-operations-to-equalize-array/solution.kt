class Solution {
    fun minOperations(nums: IntArray): Int {
        return if (nums.none { it != nums[0] }) 0 else 1
    }
}