class Solution {
    fun absDifference(nums: IntArray, k: Int): Int {
        nums.sort()
        val sumSmallest = nums.take(k).sum()
        val sumLargest = nums.takeLast(k).sum()
        return sumLargest - sumSmallest
    }
}