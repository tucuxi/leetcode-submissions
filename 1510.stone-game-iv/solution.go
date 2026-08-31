func winnerSquareGame(n int) bool {
    var sq []int

    for i := 1; i*i <= n; i++ {
        sq = append(sq, i*i)
    }

    dp := make([]bool, n+1)

    for i := 1; i <= n; i++ {
        for j := 0; j < len(sq) && sq[j] <= i; j++ {
            if !dp[i - sq[j]] {
                dp[i]= true
                break
            }
        }
    }

    return dp[n]
}