class Solution {
    fun thirdMax(nums: IntArray): Int {
        val sortedDistinctNums = nums.sorted().distinct()
        return if (sortedDistinctNums.size < 3)
            sortedDistinctNums.last()
        else
            sortedDistinctNums[sortedDistinctNums.size - 3]
    }
}