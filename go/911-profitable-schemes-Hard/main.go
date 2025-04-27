const MOD = 1_000_000_007

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
    dp := make([][][]int, len(group) + 1)
    for i := range dp {
        dp[i] = make([][]int, n + 1)
        for j := range dp[i] {
            dp[i][j] = make([]int, minProfit + 1)
        }
    }
    dp[0][0][0] = 1
    for i := 1; i < len(dp); i++ {
        g := group[i - 1]
		p := profit[i - 1]
        for j := range dp[i] {
            for k := range dp[i][j] {
                if j >= g {
                    dp[i][j][k] = (dp[i - 1][j][k] + dp[i - 1][j - g][max(0, k - p)]) % MOD
                } else {
                    dp[i][j][k] = dp[i - 1][j][k]
                } 
            }
        }
    }
    res := 0
    for j := 0; j <= n; j++ {
        res = (res + dp[len(group)][j][minProfit]) % MOD
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}