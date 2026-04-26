class Solution {
    fun maxDistance(nums1: IntArray, nums2: IntArray): Int {
        var res = 0
        var i = 0
        var j = 0

        while (i < nums1.size && j < nums2.size) {
            if (nums1[i] <= nums2[j]) {
                res = maxOf(res, j - i)
                j++
            } else {
                i++
            }
        }

        return res
    }
}