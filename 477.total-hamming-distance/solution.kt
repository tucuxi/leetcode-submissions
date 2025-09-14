class Solution {
    fun totalHammingDistance(nums: IntArray): Int {
        return (0..29).sumOf { bit ->
            val mask = 1 shl bit
            val ones = nums.count { it and mask != 0 }
            ones * (nums.size - ones)
        }
    }
}