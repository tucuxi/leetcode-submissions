class Solution {
    fun minimumPrefixLength(nums: IntArray): Int {
        var i = nums.lastIndex
        while (i > 0 && nums[i - 1] < nums[i]) {
            i--
        }
        return i
    }
}