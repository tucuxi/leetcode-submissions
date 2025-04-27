func buyChoco(prices []int, money int) int {
    min1, min2 := math.MaxInt, math.MaxInt
    for _, p := range prices {
        if p < min1 {
            min1, min2 = p, min1
        } else if p < min2 {
            min2 = p
        }
    }
    res := money - min1 - min2
    if res < 0 {
        return money
    }
    return res
}