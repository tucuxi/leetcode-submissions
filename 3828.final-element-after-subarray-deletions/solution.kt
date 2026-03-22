class Solution {
    fun finalElement(nums: IntArray): Int {
        return maxOf(nums.first(), nums.last())
    }
}