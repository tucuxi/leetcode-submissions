func minPrice(prices []int, discounts []int) float64 {
    slices.Sort(prices)
    slices.Sort(discounts)

    res := 0.
    j := len(discounts) - 1

    for i := len(prices) - 1; i >= 0; i-- {
        if j >= 0 {
            f := float64(100 - discounts[j]) / 100
            res += float64(prices[i]) * f
            j--
        } else {
            res += float64(prices[i])
        }  
    }

    return res
}