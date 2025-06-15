const mod = 1_000_000_007

func numWays(words []string, target string) int {
    k := len(words[0])
    m := len(target)
    h := make([][26]int, k)
    for _, w := range words {
        for i := range w {
            h[i][w[i] - 'a']++
        }
    }
    dp := make([][]int64, m + 1)
    for i := range dp {
        dp[i] = make([]int64, k + 1)
    }
    dp[0][0] = 1
    for i := 0; i <= m; i++ {
        for j := 0; j < k; j++ {
            if i < m {
                dp[i + 1][j + 1] = dp[i][j] * int64(h[j][target[i] - 'a']) % mod
            }
            dp[i][j + 1] += dp[i][j]
            dp[i][j + 1] %= mod
        }
    }
    return int(dp[m][k])
}