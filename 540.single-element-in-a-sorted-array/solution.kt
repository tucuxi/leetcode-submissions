class Solution {
    fun singleNonDuplicate(nums: IntArray): Int {
        var low = 0
        var high = nums.lastIndex
        while (low < high) {
            val mid = (low + (high - low) / 2) and 1.inv()
            if (nums[mid] == nums[mid + 1]) {
                low = mid + 2
            } else {
                high = mid - 1
            }
        }
        return nums[low]
    }
}