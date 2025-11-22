class Solution {
    fun kLengthApart(nums: IntArray, k: Int): Boolean {
        return nums
            .withIndex()
            .filter { it.value == 1 }
            .zipWithNext()
            .all { it.second.index - it.first.index > k }
    }
}