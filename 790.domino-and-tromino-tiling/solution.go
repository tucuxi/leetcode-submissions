const mod = 1_000_000_007

func numTilings(n int) int {
    if n <= 2 {
        return n
    }
    dp0, dp1, dp2 := 1, 1, 2
    for i := 3; i <= n; i++ {
        dp0, dp1, dp2 = dp1, dp2, (2 * dp2 + dp0) % mod
    }
    return dp2
}