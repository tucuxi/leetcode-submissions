class Solution {
    fun findDifferentBinaryString(nums: Array<String>): String {
        return buildString {
            for (i in nums.indices) {
                append('1' - nums[i][i])
            }
        }
    }
}