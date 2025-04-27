const mod = 1_000_000_007

func countRoutes(locations []int, start int, finish int, fuel int) int {
    dp := make([][]int, len(locations))
    for i := range dp {
        dp[i] = make([]int, fuel + 1)
    }
    for k := 0; k <= fuel; k++ {
        dp[finish][k] = 1
        for i := range locations {
            for j := range locations {
                if i == j {
                    continue
                }
                if f := abs(locations[i] - locations[j]); k >= f {
                    dp[i][k] = (dp[i][k] + dp[j][k - f]) % mod
                }
            }
        }
    }
    return dp[start][fuel]
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}