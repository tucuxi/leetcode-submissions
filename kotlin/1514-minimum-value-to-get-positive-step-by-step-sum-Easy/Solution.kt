class Solution {
    fun minStartValue(nums: IntArray): Int {
        var prefixSum = 0
        var minPrefixSum = Int.MAX_VALUE
        for (n in nums) {
            prefixSum += n
            minPrefixSum = minOf(minPrefixSum, prefixSum)
        }
        return maxOf(1, 1 - minPrefixSum)
    }
}