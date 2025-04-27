class Solution {
    fun nextGreaterElement(nums1: IntArray, nums2: IntArray): IntArray {
        val res = IntArray(nums1.size)
        for (i1 in nums1.indices) {
            var i2 = 0
            while (i2 <= nums2.lastIndex && nums2[i2] != nums1[i1]) {
                i2++
            }
            while (i2 <= nums2.lastIndex && nums2[i2] <= nums1[i1]) {
                i2++
            }
            res[i1] = if (i2 <= nums2.lastIndex) nums2[i2] else -1
        }
        return res
    }
}