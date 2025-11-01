class Solution {
    fun maxFrequency(nums: IntArray, k: Int, numOperations: Int): Int {
        val h = nums.asSequence().groupingBy { it }.eachCount()
        val n = h.keys.max() + k + 1
        
        val prefix = IntArray(n + 1)
        for (i in 0 until n) {
            prefix[i + 1] = prefix[i] + h.getOrDefault(i, 0)
        }

        var res = 0
        for (i in 1 until n) {
            var left = maxOf(0, i - k)
            var right = minOf(n - 1, i + k)
            val total = prefix[right + 1] - prefix[left]
            val add = total - h.getOrDefault(i, 0)
            val freq = h.getOrDefault(i, 0) + minOf(numOperations, add)
            res = maxOf(res, freq)
        }
        
        return res
    }
}