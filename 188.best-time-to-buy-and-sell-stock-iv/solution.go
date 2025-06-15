func maxProfit(k int, prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, len(prices))
        if i == 0 {
            continue
        }
        for j := 1; j < len(dp[i]); j++ {
            m := 0
            for p := 0; p < j; p++ {
                s := prices[j] - prices[p] + dp[i-1][p]
                if s > m {
                    m = s
                }
            }
            if m > dp[i][j-1] {
                dp[i][j] = m
            } else {
                dp[i][j] = dp[i][j-1]
            }
        }
    }
    return dp[k][len(prices)-1]
}