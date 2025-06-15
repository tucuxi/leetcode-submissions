func kthCharacter(k int) byte {
    dp := make([]byte, 0, k + 1)
    dp = append(dp, 0, 'a', 'b')
    p := 1
    q := 2
    i := 0
    for len(dp) <= k {
        if i == q {
            p = 1
            q *= 2
            i = 0
        }
        r := dp[p]
        if r == 'z' {
            r = 'a'
        } else {
            r++
        }
        dp = append(dp, r)
        p++
        i++
    }
    return dp[k]
}