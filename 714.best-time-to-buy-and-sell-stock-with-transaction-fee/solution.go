func maxProfit(prices []int, fee int) int {
    free, hold := 0, -prices[0]
    for _, p := range prices[1:] {
        free, hold = max(free, hold + p - fee), max(hold, free - p)
    }
    return free
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}