func maxProfit(prices []int) int {
    p := 0
    h := prices[0]
    l := len(prices) - 1
    for i := 1; i <= l; i++ {
        if prices[i] < prices[i-1] {
            p += prices[i-1] - h
            h = prices[i]
        }
    }
    if h < prices[l] {
        p += prices[l] - h
    }
    return p
}
