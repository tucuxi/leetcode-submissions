class Solution {
    fun maxProductDifference(nums: IntArray): Int =
        nums.sorted().let { it.last() * it[it.lastIndex - 1] - it.first() * it[1] }
}