func cherryPickup(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    dp := make([][][]int, rows+1)
    for i := range dp {
        dp[i] = make([][]int, cols)
        for j := range dp[i] {
            dp[i][j] = make([]int, cols)
        }
    }
    for i := rows-1; i >= 0; i-- {
        for j := 0; j < cols; j++ {
            for k := 0; k < cols; k++ {
                dp[i][j][k] = best(dp, i+1, j, k) + grid[i][j]
                if j != k {
                    dp[i][j][k] += grid[i][k]
                }
            }
        }
    }
    return dp[0][0][cols-1]
}

func best(dp [][][]int, row, col1, col2 int) int {
    res := 0
    for i := max(0, col1-1); i <= min(len(dp[row]) - 1, col1+1); i++ {
        for j := max(0, col2-1); j <= min(len(dp[row]) - 1, col2+1); j++ {
            res = max(res, dp[row][i][j])
        }
    }
    return res
} 