func minDistance(word1 string, word2 string) int {
    l := longestCommonSubsequence(word1, word2)
    return len(word1) - l + len(word2) - l
}

func longestCommonSubsequence(text1 string, text2 string) int {
    dp := make([][]int, len(text1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(text2) + 1)
    }
    for i := 0; i < len(text1); i++ {
        for j := 0; j < len(text2); j++ {
            if text1[i] == text2[j] {
                dp[i + 1][j + 1] = dp[i][j] + 1
            } else {
                dp[i + 1][j + 1] = max(dp[i][j + 1], dp[i + 1][j])
            } 
        }
    }
    return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}