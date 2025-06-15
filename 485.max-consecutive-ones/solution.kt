class Solution {
    fun findMaxConsecutiveOnes(nums: IntArray): Int {
        var currentLen = 0
        var maxLen = 0
        for (num in nums) {
            currentLen = when (num) {
                1 -> currentLen + 1
                else -> 0
            }
            if (currentLen > maxLen)
                maxLen = currentLen
        }
        return maxLen
    }
}