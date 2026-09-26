func distinctSubseqII(s string) int {
    const MOD = 1_000_000_007
    var last [26]int

    for i := range last {
        last[i] = -1
    }

    n := len(s)
    dp := make([]int, n+1)
    dp[0] = 1

    for i := range s {
        x := s[i] - 'a'
        dp[i+1] = 2 * dp[i] % MOD
        if l := last[x]; l >= 0 {
            dp[i+1] -= dp[l]
        }
        dp[i+1] %= MOD
        last[x] = i
    }

    dp[n]--
    if dp[n] < 0 {
        dp[n] += MOD
    }
    return dp[n]
}