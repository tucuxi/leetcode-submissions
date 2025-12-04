class Solution {
    fun maxSubarraySum(nums: IntArray, k: Int): Long {
        var prefixSum = 0L
        var maxSum = Long.MIN_VALUE
        val kSum = LongArray(k) { Long.MAX_VALUE / 2 }
        kSum[k - 1] = 0

        nums.forEachIndexed { i, num ->
            prefixSum += num
            maxSum = maxOf(maxSum, prefixSum - kSum[i % k])
            kSum[i % k] = minOf(kSum[i % k], prefixSum)
        }

        return maxSum
    }
}