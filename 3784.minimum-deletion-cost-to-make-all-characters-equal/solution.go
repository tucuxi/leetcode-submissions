func minCost(s string, cost []int) int64 {
    var (
        total    int64
        charCost [26]int64
        maxCost  int64
    )
    for i := range s {
        ci := s[i] - 'a'
        total += int64(cost[i])
        charCost[ci] += int64(cost[i])      
        maxCost = max(maxCost, charCost[ci])
    }
    return total - maxCost
}