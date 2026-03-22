class Solution {
    fun minCost(nums1: IntArray, nums2: IntArray): Int {
        val counts = mutableMapOf<Int, Int>()
        
        nums1.forEach { counts[it] = counts.getOrDefault(it, 0) + 1 }
        nums2.forEach { counts[it] = counts.getOrDefault(it, 0) - 1 }

        val v = counts.values

        return if (v.any { it % 2 != 0 }) {
            -1
        } else {
            v.filter { it > 0 }.sum() / 2
        }
    }
}