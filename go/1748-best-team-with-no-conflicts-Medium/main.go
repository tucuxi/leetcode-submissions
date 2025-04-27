type player struct {
    age int
    score int
}

func bestTeamScore(scores []int, ages []int) int {
    res := 0
    n := len(scores)
    players := make([]player, n)
    for i := range players {
        players[i] = player{age: ages[i], score: scores[i]}
    }
    sort.Slice(players, func(i, j int) bool {
        return players[i].age < players[j].age ||
            players[i].age == players[j].age && players[i].score < players[j].score
    })
    dp := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = players[i].score
        for j := i - 1; j >= 0; j-- {
            if players[i].score >= players[j].score {
                if s := players[i].score + dp[j]; s > dp[i] {
                    dp[i] = s
                } 
            }
        }
        if dp[i] > res {
            res = dp[i]
        }
    }
    return res
}