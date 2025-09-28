class Solution {
    fun maxKDistinct(nums: IntArray, k: Int): IntArray {
        return nums.distinct().sortedDescending().take(k).toIntArray()
    }
}