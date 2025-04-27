func stoneGameIII(stoneValue []int) string {
    memo := make([]int, len(stoneValue))
    for i := range memo {
        memo[i] = math.MinInt
    }
    
    var dp func(int) int
    dp = func(i int) int {
        if i == len(stoneValue) {
            return 0
        }
        if memo[i] != math.MinInt {
            return memo[i]
        }
        s := 0
        for j := 0; j <= 2 && i + j < len(stoneValue); j++ {
            s += stoneValue[i + j]
            if h := s - dp(i + j + 1); h > memo[i] {
                memo[i] = h
            }
        }
        return memo[i]
    }
    
    switch d0 := dp(0); {
    case d0 > 0:
        return "Alice"
    case d0 < 0:
        return "Bob"
    default:
        return "Tie"
    }
}