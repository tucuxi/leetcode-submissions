func ways(pizza []string, k int) int {
    rows := len(pizza)
    cols := len(pizza[0])
    apples := make([][]int, rows)
    for i := rows - 1; i >= 0; i-- {
        apples[i] = make([]int, cols)
        for j := cols - 1; j >= 0; j-- {
            if pizza[i][j] == 'A' {
                apples[i][j] = 1
            } 
            if i + 1 < rows {
                apples[i][j] += apples[i + 1][j]
            }
            if j + 1 < cols {
                apples[i][j] += apples[i][j + 1]
            }
            if i + 1 < rows && j + 1 < cols {
                apples[i][j] -= apples[i + 1][j + 1]
            }
        }
    }
    dp := make([][][]int, rows)
    for i := 0; i < rows; i++ {
        dp[i] = make([][]int, cols)
        for j := 0; j < cols; j++ {
            dp[i][j] = make([]int, k)
            if apples[i][j] > 0 {
                dp[i][j][0] = 1
            }
        }
    }
    for l := 1; l < k; l++ {
        for i := rows - 1; i >= 0; i-- {
            for j := cols - 1; j >= 0; j-- {
                for p := i; p < rows - 1; p++ {
                    if apples[i][j] - apples[p + 1][j] > 0 {
                        dp[i][j][l] = (dp[i][j][l] + dp[p + 1][j][l - 1]) % 1_000_000_007;
                    }
                }
                for q := j; q < cols - 1; q++ {
                    if apples[i][j] - apples[i][q + 1] > 0 {
                        dp[i][j][l] = (dp[i][j][l] + dp[i][q + 1][l - 1]) % 1_000_000_007;
                    }
                }
            }
        }
    }
    return dp[0][0][k - 1];
}