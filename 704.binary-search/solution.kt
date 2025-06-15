class Solution {
    fun search(nums: IntArray, target: Int): Int {
        var lo = 0
        var hi = nums.size - 1
        do {
            val m = lo + (hi - lo) / 2
            if (nums[m] < target)
                lo = m + 1
            else if (nums[m] > target)
                hi = m - 1
            else
                return m
        } while (lo <= hi)
        return -1
    }
}