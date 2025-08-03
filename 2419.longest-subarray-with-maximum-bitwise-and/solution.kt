class Solution {
    fun longestSubarray(nums: IntArray): Int {
        val (maxLength, _, _) = nums.fold(Triple(0, 0, 0)) { (maxLength, curLength, maxBitwise), num ->
            when {
                num > maxBitwise -> Triple(1, 1, num)
                num == maxBitwise -> Triple(maxOf(curLength + 1, maxLength), curLength + 1, maxBitwise)
                else -> Triple(maxLength, 0, maxBitwise)
            }
        }
        return maxLength
    }
}