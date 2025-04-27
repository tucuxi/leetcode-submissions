func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    if len(s1) < len(s2) {
        return isInterleave(s2, s1, s3)
    }
    dp := make([][]bool, len(s1) + 1)
    for i := range dp {
        dp[i] = make([]bool, len(s2) + 1)
    }
    dp[0][0] = true
    for i := 1; i <= len(s1); i++ {
        dp[i][0] = dp[i - 1][0] && s1[i - 1] == s3[i - 1]
    }
    for j := 1; j <= len(s2); j++ {
        dp[0][j] = dp[0][j - 1] && s2[j - 1] == s3[j - 1]
    }
    for i := 1; i <= len(s1); i++ {
        for j := 1; j <= len(s2); j++ {
            dp[i][j] = dp[i - 1][j] && s1[i - 1] == s3[i + j - 1] || dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]
        }
    }
    return dp[len(s1)][len(s2)]
}