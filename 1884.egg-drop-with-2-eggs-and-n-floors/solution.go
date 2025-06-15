func twoEggDrop(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = i
        for j := 1; j < i; j++ {
            dp[i] = min(dp[i], max(j, dp[i-j] + 1))
        }
    }
    return dp[n]
}