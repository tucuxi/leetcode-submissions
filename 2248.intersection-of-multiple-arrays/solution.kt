class Solution {
    fun intersection(nums: Array<IntArray>): List<Int> {
        return nums.fold(nums[0].toSet()) { acc, elem -> elem intersect acc }.toList().sorted()
    }
}