class Solution {
    fun findDifference(nums1: IntArray, nums2: IntArray): List<List<Int>> {
        return listOf(
            nums1.filter { it !in nums2 }.distinct(),
            nums2.filter { it !in nums1 }.distinct()
        )
    }
}