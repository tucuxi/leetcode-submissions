class Solution {
    fun maximumScore(nums: IntArray): Long {
        val suffixMin = IntArray(nums.size - 1)
        suffixMin[suffixMin.lastIndex] = nums[nums.lastIndex]
        for (i in suffixMin.lastIndex - 1 downTo 0) {
            suffixMin[i] = minOf(nums[i + 1], suffixMin[i + 1])
        }

        var maxScore = Long.MIN_VALUE
        var prefixSum = 0L
        for (i in 0 until nums.lastIndex) {
            prefixSum += nums[i].toLong()
            val score = prefixSum - suffixMin[i].toLong()
            maxScore = maxOf(maxScore, score)
        } 

        return maxScore
    }
}