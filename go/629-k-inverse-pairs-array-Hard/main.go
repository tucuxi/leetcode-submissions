const mod = 1_000_000_007

func kInversePairs(n int, k int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        dp[i][0] = 1
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= k; j++ {
            dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
            if j >= i {
                dp[i][j] = (dp[i][j] - dp[i-1][j-i] + mod) % mod
            }
        }
    }
    return dp[n][k]
}