class Solution {
    fun shuffle(nums: IntArray, n: Int): IntArray {
        return sequence {
            for (i in 0 until n) {
                yield(nums[i])
                yield(nums[i + n])
            }
        }.toList().toIntArray()
    }
}