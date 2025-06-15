const M = 1_000_000_007

func numberOfArrays(s string, k int) int {
    dp := make([]int, len(s))
    for i := range dp {
        dp[i] = -1
    }
    
    var numberOfArraysMemo func(int) int
    numberOfArraysMemo = func(i int) int {
        if i == len(s) {
            return 1
        }
        if s[i] == '0' {
            return 0
        }
        if dp[i] >= 0 {
            return dp[i]
        }
        dp[i] = 0
        for j := i + 1; j <= len(s); j++ {
            if n, _ := strconv.Atoi(s[i:j]); n > k {
                break
            }
            dp[i] = (dp[i] + numberOfArraysMemo(j)) % M
        }
        return dp[i]
    }
    
    return numberOfArraysMemo(0)
}