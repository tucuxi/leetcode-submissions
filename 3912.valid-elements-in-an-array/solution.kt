class Solution {
    fun findValidElements(nums: IntArray): List<Int> {
        val rightMax = IntArray(nums.size)
        rightMax[nums.lastIndex] = 0

        for (i in nums.lastIndex - 1 downTo 0) {
            rightMax[i] = maxOf(nums[i + 1], rightMax[i + 1])
        }

        var leftMax = 0
        val res = mutableListOf<Int>()

        for (i in nums.indices) {
            if (nums[i] > leftMax || nums[i] > rightMax[i]) {
                res += nums[i]
            }
            leftMax = maxOf(leftMax, nums[i])
        }

        return res
    }
}