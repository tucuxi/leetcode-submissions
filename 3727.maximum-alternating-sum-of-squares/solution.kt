class Solution {
    fun maxAlternatingSum(nums: IntArray): Long {
        for (i in nums.indices) {
            nums[i] *= nums[i]
        }
        
        nums.sort()

        val n = nums.size
        val p = (n / 2 until n).sumOf { nums[it].toLong() }
        val q = (0 until n / 2).sumOf { nums[it].toLong() }

        return p - q
    }
}