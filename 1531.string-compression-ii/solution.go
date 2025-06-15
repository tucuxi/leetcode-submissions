func getLengthOfOptimalCompression(s string, k int) int {
    n := len(s)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        for j := range dp[i] {
            dp[i][j] = n
        }
    }
    dp[0][0] = 0
    for i := 1; i <= n; i++ {
        for j := 0; j <= k; j++ {
            if j > 0 {
                dp[i][j] = dp[i-1][j-1]
            }
            removed, cnt := 0, 0
            for p := i; p > 0; p-- {
                if s[p-1] == s[i-1] { 
                    cnt++ 
                } else {
                    removed++
                    if removed > j {
                        break
                    }
                }
                dp[i][j] = min(dp[i][j], dp[p-1][j-removed] + enc(cnt))
            }
        }
    }
    return dp[n][k]
}

func enc(len int) int {
    switch  {
        case len == 0: return 0
        case len == 1: return 1
        case len < 10: return 2
        case len < 100: return 3
        default: return 4
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}