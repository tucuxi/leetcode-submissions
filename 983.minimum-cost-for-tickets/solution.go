func mincostTickets(days []int, costs []int) int {
    lastDay := days[len(days) - 1]
    dp := make([]int, lastDay + 1)
    k := 0
    for i := 1; i <= lastDay; i++ {
        if i < days[k] {
            dp[i] = dp[i - 1]
        } else {
            dp[i] = min(dp[i - 1] + costs[0], dp[max(0, i - 7)] + costs[1], dp[max(0, i - 30)] + costs[2]) 
            k++
        }
    }
    return dp[lastDay]
}