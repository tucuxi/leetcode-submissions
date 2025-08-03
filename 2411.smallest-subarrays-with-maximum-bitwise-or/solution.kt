class Solution {
    fun smallestSubarrays(nums: IntArray): IntArray {
        val res = IntArray(nums.size)
        val closestRightIndices = IntArray(32) { -1 }
        nums.withIndex().reversed().forEach { (index, num) ->
            var rightIndex = index
            (0 until 32).forEach { bit ->
                if (num and (1 shl bit) == 0) {
                    rightIndex = maxOf(rightIndex, closestRightIndices[bit])
                } else {
                    closestRightIndices[bit] = index
                }
            }
            res[index] = rightIndex - index + 1
        }
        return res
    }
}