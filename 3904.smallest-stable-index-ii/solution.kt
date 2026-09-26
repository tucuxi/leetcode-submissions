class Solution {
    fun firstStableIndex(nums: IntArray, k: Int): Int {
        val n = nums.size
        val mn = IntArray(n)
        
        mn[n - 1] = nums[n - 1]
        for (i in n - 2 downTo 0) {
            mn[i] = minOf(nums[i], mn[i + 1])
        }

        var mx = nums[0]
        for (i in 0 until n) {
            mx = maxOf(nums[i], mx)
            if (mx - mn[i] <= k) {
                return i
            }
        }
        return -1
    }
}