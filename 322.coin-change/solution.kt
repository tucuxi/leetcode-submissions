class Solution {
    fun coinChange(coins: IntArray, amount: Int): Int {
        val dp = IntArray(amount + 1) { if (it == 0) 0 else amount + 1 }
        for (i in 1..amount) {
            for (c in coins) {
                if (c <= i) {
                    dp[i] = minOf(dp[i], dp[i - c] + 1)
                } 
            }
        }
        return if (dp[amount] > amount) -1 else dp[amount]
    }
}