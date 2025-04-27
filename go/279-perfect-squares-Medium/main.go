func numSquares(n int) int {
    dp := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        dp[i] = n + 1
        for j := 1; j * j <= i; j++ {
            if k := i - j * j; dp[k] < dp[i] {
                dp[i] = dp[k] + 1
            }
        }
    }
    return dp[n]
}