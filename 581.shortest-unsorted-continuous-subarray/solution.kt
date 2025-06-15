class Solution {
    fun findUnsortedSubarray(nums: IntArray): Int {
        return nums
            .sortedArray()
            .zip(nums)
            .dropWhile { (a, b) -> a == b }
            .dropLastWhile { (a, b) -> a == b }
            .size
    }
}