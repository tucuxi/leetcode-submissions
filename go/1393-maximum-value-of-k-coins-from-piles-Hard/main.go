func maxValueOfCoins(piles [][]int, k int) int {
    n := len(piles)
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k + 1)
    }
    for i := 1; i <= n; i++ {
        for j := 0; j <= k; j++ {
            s := 0
            for l := 0; l <= j && l <= len(piles[i - 1]); l++ {
                if l > 0 {
                    s += piles[i - 1][l - 1]
                }
                if h := dp[i - 1][j - l] + s; h > dp[i][j] {
                    dp[i][j] = h
                }
            }
        }
    }
    return dp[n][k]
}