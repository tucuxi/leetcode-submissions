func countSquares(matrix [][]int) int {
    res := 0
    n := len(matrix[0])
    dp := make([][]int, len(matrix))
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = matrix[i][0]
        res += dp[i][0]
    }
    for j := 1; j < n; j++ {
        dp[0][j] = matrix[0][j]
        res += dp[0][j]
    }
    for i := 1; i < len(dp); i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 1 {
                dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1 
                res += dp[i][j]
            }
        }
    }
    return res
}