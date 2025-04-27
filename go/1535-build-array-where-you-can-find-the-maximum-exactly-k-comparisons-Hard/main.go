func numOfArrays(n int, m int, k int) int {
    const mod = 1_000_000_007
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, m + 1)
        for j := range dp[i] {
            dp[i][j] = make ([]int, k + 1)
        }
    }
    for num := range dp[0] {
        dp[n][num][0] = 1
    }
    for i := n - 1; i >= 0; i-- {
        for maxSoFar := m; maxSoFar >= 0; maxSoFar-- {
            for remain := 0; remain <= k; remain++ {
                ans := 0
                for num := 1; num <= maxSoFar; num++ {
                    ans = (ans + dp[i + 1][maxSoFar][remain]) % mod
                }
                if remain > 0 {
                    for num := maxSoFar + 1; num <= m; num++ {
                        ans = (ans + dp[i + 1][num][remain - 1]) % mod
                    }
                } 
                dp[i][maxSoFar][remain] = ans
            }
        }
    }
    return dp[0][0][k]
}