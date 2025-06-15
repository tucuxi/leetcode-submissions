class Solution {
    fun maximumValueSum(nums: IntArray, k: Int, edges: Array<IntArray>): Long {
        val sum = nums.sumOf { it.toLong() }
        val increase = nums
            .map { num -> (num xor k) - num }
            .sortedDescending()
            .chunked(2) { c -> if (c.size < 2) 0L else { maxOf(0, c.sum()).toLong() }}
            .sum()
        return sum + increase
    }
}