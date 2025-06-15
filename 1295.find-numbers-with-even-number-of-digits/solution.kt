class Solution {
    fun findNumbers(nums: IntArray): Int =
        nums.count { n -> n >= 10 && n < 100 || n >= 1000 && n < 10000 || n == 100000}
}