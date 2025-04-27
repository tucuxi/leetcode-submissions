class Solution {
    fun decompressRLElist(nums: IntArray): IntArray {
        return sequence {
            for (i in 0 until nums.size step 2) {
                for (j in 0 until nums[i])
                    yield(nums[i + 1])
            }
        }.toList().toIntArray()
    }
}