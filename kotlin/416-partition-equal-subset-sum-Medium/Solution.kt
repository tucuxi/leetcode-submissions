class Solution {
    fun canPartition(nums: IntArray): Boolean {
        val total = nums.sum()
        
        if (total % 2 == 1) {
            return false
        }

        val target = total / 2
        val subset = mutableSetOf(0)

        for (num in nums) {
            for (sum in subset.toList()) {
                val newSum = sum + num
                if (newSum == target) {
                    return true
                }
                if (newSum < target) {
                    subset += newSum
                }
            }
        }

        return false
    }
}