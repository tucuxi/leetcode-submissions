func isScramble(s1 string, s2 string) bool {
    n := len(s1)
    letters1, letters2 := []byte(s1), []byte(s2)
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n) 
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if letters1[i] == letters2[j] {
                dp[1][i][j] = 1
            }
        }
    }

    for length := 2; length <= n; length++ {
        for i := 0; i < n + 1 - length; i++ {
            for j := 0; j < n + 1 -length; j++ {
                for newLength := 1; newLength < length; newLength++ {
                    dp1 := dp[newLength][i]
                    dp2 := dp[length - newLength][i + newLength]

                    dp[length][i][j] |= dp1[j] & dp2[j + newLength]
                    dp[length][i][j] |= dp1[j + length - newLength] & dp2[j]
                }
            }
        }
    }

    return dp[n][0][0] > 0
}