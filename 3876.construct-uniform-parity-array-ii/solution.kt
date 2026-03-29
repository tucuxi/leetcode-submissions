class Solution {
    fun uniformArray(nums1: IntArray): Boolean {
        val odds = TreeSet(nums1.filter { it % 2 != 0 })
        return nums1.all { it % 2 == 0 || odds.lower(it) != null }
            || nums1.all { it % 2 != 0 || odds.lower(it) != null }
    }
}