func maximalSquare(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    maxSide := 0
    for i := range m {
        dp[i] = make([]int, n)
        for j := range n {
            if matrix[i][j] == '0' {
                continue
            }
            if i == 0 || j == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
            maxSide = max(maxSide, dp[i][j])
        }
    }
    return maxSide * maxSide
}