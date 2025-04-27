func paintWalls(cost []int, time []int) int {
    memo := make([][]int, len(cost))
    for i := range memo {
        memo[i] = make([]int, len(cost) + 1)
    }
    
    var dp func(int, int) int
    dp = func(i, remain int) int {
        if remain <= 0 {
            return 0
        }
        if i == len(cost) {
            return 1_000_000_000
        }
        if memo[i][remain] > 0 {
            return memo[i][remain]
        }
        paint := dp(i + 1, remain - time[i] - 1) + cost[i]
        dont := dp(i + 1, remain)
        if paint < dont {
            memo[i][remain] = paint
        } else {
            memo[i][remain] = dont
        }
        return memo[i][remain]
    }
    
    return dp(0, len(cost))
}