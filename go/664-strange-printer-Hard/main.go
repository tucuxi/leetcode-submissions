func strangePrinter(s string) int {
    dp := make([][]int, len(s))
    for i := range dp {
        dp[i] = make([]int, len(s))
    }
    for length := 1; length <= len(s); length++ {
        for left := 0; left <= len(s) - length; left++ {
            right := left + length - 1
            dp[left][right] = length + 1
            i := left
            for ; i < right && s[i] == s[right]; i++ {}
            if i == right {
                dp[left][right] = 0
            } else {
                for j := i; j < right; j++ {
                    if k := dp[i][j] + dp[j + 1][right] + 1; k < dp[left][right] {
                        dp[left][right] = k
                    }                
                }
            }
        }
    }
    return dp[0][len(s) - 1] + 1
}