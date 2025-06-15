func finalPrices(prices []int) []int {
    res := make([]int, len(prices))
    copy(res, prices)

    var stack []int

    for i, price := range prices {
        for len(stack) > 0 && prices[stack[len(stack) - 1]] >= price {
            res[stack[len(stack) - 1]] -= price
            stack = stack[:len(stack) - 1]
        }    
        stack = append(stack, i)
    }
    return res
}