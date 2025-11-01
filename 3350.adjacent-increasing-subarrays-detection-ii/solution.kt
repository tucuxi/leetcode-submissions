class Solution {
    fun maxIncreasingSubarrays(nums: List<Int>): Int {
        var previousLength = 0
        var currentLength = 1
        var result = 0

        for (i in 1 .. nums.lastIndex) {
            if (nums[i] <= nums[i - 1]) {
                previousLength = currentLength
                currentLength = 0
            }
            currentLength++
            result = maxOf(result, maxAdjacent(previousLength, currentLength))
        }
        return result
    }

    private fun maxAdjacent(a: Int, b: Int) = max(min(a, b), max(a, b) / 2)
}