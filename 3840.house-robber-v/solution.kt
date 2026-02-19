class Solution {
    fun rob(nums: IntArray, colors: IntArray): Long {
        if (nums.size == 0) return 0L

        var prev2 = 0L
        var prev = nums[0].toLong()

        for (i in 1 until nums.size) {
            val currentVal = nums[i].toLong()
            val current = if (colors[i] != colors[i - 1]) {
                prev + currentVal
            } else {
                maxOf(prev, currentVal + prev2)
            }
            prev2 = prev
            prev = current
        }

        return prev
    }
}