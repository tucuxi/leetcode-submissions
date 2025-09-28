class Solution {
    fun maxTotalValue(nums: IntArray, k: Int): Long {
        val d = nums.max() - nums.min()
        return d.toLong() * k
    }
}