func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := math.MaxInt
    for _, price := range prices {
        if trade := price - minPrice; trade > maxProfit {
            maxProfit = price - minPrice
        }
        if price < minPrice {
            minPrice = price
        }
    }
    return maxProfit
}