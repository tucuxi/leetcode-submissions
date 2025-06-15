func minimumTotal(triangle [][]int) int {
    dp := make([][]int, len(triangle))
    for i := range dp {
        dp[i] = make([]int, i + 1)
    }
    dp[0][0] = triangle[0][0]
    for i := 1; i < len(triangle); i++ {
        for j := 0; j <= i; j++ {
            min := math.MaxInt
            if j < len(dp[i-1]) && dp[i-1][j] < min {
                min = dp[i-1][j]
            }
            if j > 0 && dp[i-1][j-1] < min {
                min = dp[i-1][j-1]
            }
            dp[i][j] = min + triangle[i][j]
        }
    }
    lastRow := len(dp) - 1
    min := dp[lastRow][0]
    for i := 1; i < len(dp[lastRow]); i++ {
        if dp[lastRow][i] < min {
            min = dp[lastRow][i]
        }
    }
    return min
}