class Solution {
    fun minCost(nums1: IntArray, nums2: IntArray): Int {
        val counts1 = nums1.asIterable().groupingBy { it }.eachCount()
        val totalCounts = counts1.toMutableMap()

        nums2.forEach { totalCounts[it] = totalCounts.getOrDefault(it, 0) + 1 }

        var res = 0L

        totalCounts.forEach { (num, totalCount) ->
            if (totalCount % 2 != 0) {
                return -1
            }
            val targetCount = totalCount / 2
            val count1 = counts1.getOrDefault(num, 0)
            res += maxOf(0, count1 - targetCount)
        }

        return res.toInt()
    }
}