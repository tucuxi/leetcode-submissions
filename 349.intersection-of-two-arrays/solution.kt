class Solution {
    fun intersection(nums1: IntArray, nums2: IntArray): IntArray {
        val intersection = nums1.toSet() intersect nums2.toSet()
        return intersection.toIntArray()
    }
}