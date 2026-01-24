class Solution {
    fun maximumProfit(prices: IntArray, k: Int): Long {
        val n = prices.size
        val dp = Array(n) { Array(k + 1) { LongArray(3) } }
        for (j in 1 .. k) {
            dp[0][j][1] = -prices[0].toLong()
            dp[0][j][2] = prices[0].toLong()
        }
        for (i in 1 until n) {
            for (j in 1 .. k) {
                dp[i][j] = longArrayOf(
                    maxOf(
                        dp[i - 1][j][0],
                        maxOf(
                            dp[i - 1][j][1] + prices[i],
                            dp[i - 1][j][2] - prices[i]
                        )
                    ),
                    maxOf(
                        dp[i - 1][j][1],
                        dp[i - 1][j - 1][0] - prices[i]
                    ),
                    maxOf(
                        dp[i - 1][j][2],
                        dp[i - 1][j - 1][0] + prices[i]
                    )
                )
            }
        }
        return dp[n - 1][k][0]
    }
}