class Solution {
    fun isTrionic(nums: IntArray): Boolean {
        val pairs = nums.asList().zipWithNext()
        val left = pairs.dropWhile { (a, b) -> a > b }
        val middle = left.dropWhile { (a, b) -> a < b }
        val right = middle.dropWhile { (a, b) -> a > b }
        val rest = right.dropWhile { (a, b) -> a < b }
        val n1 = nums.size - 1
        return left.size == n1 && middle.size < n1 && right.size < middle.size && right.size > 0 && rest.size == 0
    }
}