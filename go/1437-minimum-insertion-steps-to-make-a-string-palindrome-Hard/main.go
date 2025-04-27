func minInsertions(s string) int {
    return len(s) - lcs(s, reverse(s))
}

func reverse(s string) string {
    var t strings.Builder
    for i := len(s) - 1; i >= 0; i-- {
        t.WriteByte(s[i])
    }
    return t.String()
}

func lcs(s, t string) int {
    m, n := len(s), len(t)
    dp := make([][]int, m + 1)
    for i := range dp {
        dp[i] = make([]int, n + 1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i - 1] == t[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1
            } else {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])
            }
        }
    }
    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}