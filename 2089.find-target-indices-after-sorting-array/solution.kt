class Solution {
    fun targetIndices(nums: IntArray, target: Int): List<Int> =
        nums
            .sorted()
            .withIndex()
            .filter { (_, value) -> value == target }
            .map { (index, _) -> index }
}