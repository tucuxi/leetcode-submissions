func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    dp[0][1] = 1
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if obstacleGrid[i - 1][j - 1] == 0 {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
            }
        }
    }
    return dp[m][n]
}