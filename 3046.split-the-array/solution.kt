class Solution {
    fun isPossibleToSplit(nums: IntArray): Boolean {
        return nums.asIterable().groupingBy { it }.eachCount().values.all { it <= 2 }
    }
}