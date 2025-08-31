class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val numToIndex = mutableMapOf<Int, Int>()
        nums.forEachIndexed { index, num ->
            val otherIndex = numToIndex.getOrElse(target - num) {
                numToIndex[num] = index
                return@forEachIndexed
            }
            return intArrayOf(otherIndex, index)
        }
        return intArrayOf()
    }
}