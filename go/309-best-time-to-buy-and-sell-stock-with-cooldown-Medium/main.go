func maxProfit(prices []int) int {
    profit1, profit2 := 0, 0
    for i := 1; i < len(prices); i++ {
        profit1, profit2 =
            max(profit1 + prices[i] - prices[i - 1], profit2),
            max(profit1, profit2)
    }
    return max(profit1, profit2)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}