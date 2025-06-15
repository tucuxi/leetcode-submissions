class Solution {
    fun searchInsert(nums: IntArray, target: Int): Int {
        var lo = 0
        var hi = nums.size - 1
        var m = 0
        while (lo <= hi) {
            m = lo + (hi - lo) / 2
            if (nums[m] < target)
                lo = m + 1
            else if (nums[m] > target)
                hi = m - 1
            else
                return m
        }
        return if (target > nums[m]) m + 1 else m
    }
}