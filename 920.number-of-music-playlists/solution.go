func numMusicPlaylists(n int, goal int, k int) int {
    const mod = 1_000_000_007
    
    dp := make([][]int, goal + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    dp[0][0] = 1
    
    for i := 1; i <= goal; i++ {
        for j := 1; j <= min(i, n); j++ {
            dp[i][j] = dp[i - 1][j - 1] * (n - j + 1) % mod
            if j > k {
                dp[i][j] = (dp[i][j] + dp[i - 1][j] * (j - k)) % mod
            }
        }
    }
    return dp[goal][n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}