func isWinner(player1 []int, player2 []int) int {
    p1, p2 := score(player1), score(player2)
    if p1 > p2 {
        return 1
    }
    if p1 < p2 {
        return 2
    }
    return 0
}

func score(pins []int) int {
    t1, t2 := false, false
    res := 0
    for _, p := range pins {
        res += p
        if t1 || t2 {
            res += p
        }
        t2 = t1
        t1 = p == 10
    }
    return res
}