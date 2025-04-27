class Solution {
    fun jump(nums: IntArray): Int {
        val k = nums.size - 1
        val dp = IntArray(nums.size)
        for (i in 1..k) {
            dp[i] = Int.MAX_VALUE
            for (j in 0 until i) {
                if (j + nums[j] >= i) {
                    dp[i] = Math.min(dp[i], dp[j] + 1)
                }
            }
        }
        return dp[k]
    }
}