class Solution {
    fun compareBitonicSums(nums: IntArray): Int {
        var sum = nums[0].toLong()
        var i = 0

        while (i < nums.lastIndex && nums[i + 1] > nums[i]) {
            sum += nums[++i]
        }

        while (i <= nums.lastIndex) {
            sum -= nums[i]
            i++
        }

        return when {
            sum > 0 -> 0
            sum < 0 -> 1
            else -> -1
        }
    }
}