class Solution {
    fun divideArray(nums: IntArray): Boolean =
        nums
            .toList()
            .groupingBy { it }
            .eachCount()
            .values
            .all { it % 2 == 0}
}