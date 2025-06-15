func minCostClimbingStairs(cost []int) int {
    h1, h2 := 0, 0
    for i := 2; i <= len(cost); i++ {
        h1, h2 = min(h1 + cost[i-1], h2 + cost[i-2]), h1
    }
    return h1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}