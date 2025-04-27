class Solution {
    fun waysToSplitArray(nums: IntArray): Int {
        val prefixSum = nums.runningFold(0L) { acc, num -> acc + num }
        val total = prefixSum[prefixSum.lastIndex]
        return prefixSum.slice(1 until prefixSum.lastIndex).count { it >= total - it }
    }
}