func numRollsToTarget(n int, k int, target int) int {
    const mod = 1_000_000_007
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, target + 1)
    }
    dp[0][0] = 1
    for i := 1; i <= n; i++ {
        for j := 1; j <= target; j++ {
            for l := 1; l <= min(j, k); l++ {
                dp[i][j] = (dp[i][j] + dp[i - 1][j - l]) % mod
            }
        }
    }
    return dp[n][target]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}