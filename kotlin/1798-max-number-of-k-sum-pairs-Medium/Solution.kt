class Solution {
    fun maxOperations(nums: IntArray, k: Int): Int {
        val h = nums
            .asSequence()
            .groupingBy { it }
            .eachCount()
        return h
            .map { (num, count) -> minOf(count, h.getOrDefault(k - num, 0)) }
            .sum() / 2
    }
}