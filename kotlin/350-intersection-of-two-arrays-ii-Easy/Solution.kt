class Solution {
    fun intersect(nums1: IntArray, nums2: IntArray): IntArray {
        val frequencies1 = nums1.toList().groupingBy { it }.eachCount()
        val frequencies2 = nums2.toList().groupingBy { it }.eachCount()
        val numsInCommon = frequencies1.keys intersect frequencies2.keys
        return numsInCommon.flatMap { num ->
            val count1 = frequencies1[num]!!
            val count2 = frequencies2[num]!!
            List<Int>(minOf(count1, count2)) { num }
        }.toIntArray()
    }
}