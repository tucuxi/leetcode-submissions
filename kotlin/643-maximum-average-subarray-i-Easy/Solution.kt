import kotlin.math.max

class Solution {
    fun findMaxAverage(nums: IntArray, k: Int): Double {
        var sum = nums.sliceArray(0 until k).sum()
        var maxSum = sum
        for (i in k until nums.size) {
            sum = sum - nums[i - k] + nums[i]
            maxSum = max(sum, maxSum)
        }
        return maxSum.toDouble() / k
    }
}