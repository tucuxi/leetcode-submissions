class Solution {
    fun heightChecker(heights: IntArray): Int =
        heights.zip(heights.sorted()).count { it.first != it.second }
}