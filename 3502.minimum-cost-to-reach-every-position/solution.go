func minCosts(cost []int) []int {
    res := make([]int, len(cost))
    res[0] = cost[0]
    for i := 1; i < len(cost); i++ {
        res[i] = min(cost[i], res[i - 1])
    }
    return res
}