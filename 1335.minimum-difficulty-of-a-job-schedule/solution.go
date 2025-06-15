func minDifficulty(jobDifficulty []int, d int) int {
    dp := make([][]int, len(jobDifficulty) + 1)
    for i := range dp {
        dp[i] = make([]int, d + 1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt
        }
    }
    dp[0][0] = 0
    for i := 1; i < len(dp); i++ {
        maxI := jobDifficulty[i - 1]
        for j := i - 1; j >= 0; j-- {
            for day := 1; day <= d; day++ {
                if dp[j][day - 1] != math.MaxInt {
                    if dp[j][day - 1] + maxI < dp[i][day] {
                        dp[i][day] = dp[j][day - 1] + maxI
                    }
                }
            }
            if j > 0 && jobDifficulty[j - 1] > maxI {
                maxI = jobDifficulty[j - 1]
            }
        }
    }
    if dp[len(jobDifficulty)][d] == math.MaxInt {
        return -1
    }
    return dp[len(jobDifficulty)][d]
}