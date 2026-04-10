const k = 3

func maximumAmount(coins [][]int) int {
    n := len(coins[0])
    dp := make([][k]int, n+1)
    
    for i := range dp {
        for j := range k {
            dp[i][j] = math.MinInt / 4
        }
    }
    for j := range k {
        dp[1][j] = 0
    }
    for _, c := range coins {
        for j := 1; j <= n; j++ {
            x := c[j-1]           
            dp[j][2] = max(dp[j-1][2] + x, dp[j][2] + x, dp[j-1][1], dp[j][1])
            dp[j][1] = max(dp[j-1][1] + x, dp[j][1] + x, dp[j-1][0], dp[j][0])
            dp[j][0] = max(dp[j-1][0], dp[j][0]) + x
        }
    }

    return dp[n][2]
}