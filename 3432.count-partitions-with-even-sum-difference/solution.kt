class Solution {
    fun countPartitions(nums: IntArray): Int {
        return if (nums.sum() % 2 == 0) {
            nums.size - 1
        } else {
            0
        }
    }
}