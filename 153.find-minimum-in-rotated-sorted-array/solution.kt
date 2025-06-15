class Solution {
    fun findMin(nums: IntArray): Int {
        var l = 0
        var r = nums.lastIndex
        while (l < r) {
            val m = l + (r - l) / 2
            if (nums[m] > nums[r]) {
                l = m + 1
            } else {
                r = m
            }
        }
        return nums[l]
    }
}