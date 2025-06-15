func mostPoints(questions [][]int) int64 {
    memo := make([]int64, len(questions))
    for i := range memo {
        memo[i] = -1
    }
    
    var dp func(int) int64
    dp = func(index int) int64 {
        if index >= len(questions) {
            return 0
        }
        if memo[index] < 0 {
            solve := int64(questions[index][0]) + dp(index + questions[index][1] + 1)
            skip := dp(index + 1)
            memo[index] = max(solve, skip)
        }
        return memo[index]
    }
    
    return dp(0)
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}