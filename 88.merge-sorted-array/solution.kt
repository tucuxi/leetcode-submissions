class Solution {
    fun merge(nums1: IntArray, m: Int, nums2: IntArray, n: Int): Unit {
        var i = m - 1
        var j = n - 1
        var k = m + n - 1
        merge@ for (k in m + n - 1 downTo 0) {
            nums1[k] = when {
                i >= 0 && j >= 0 -> if (nums1[i] < nums2[j]) {
                        nums2[j--]
                    } else {
                        nums1[i--]
                    }
                j >= 0 -> nums2[j--]
                else -> break@merge
            }
        }
    }
}