func maximumTastiness(price []int, k int) int {
    sort.Ints(price)
    return sort.Search(1e9, func(x int) bool {
        c := 1
        p := price[0]
        for _, q := range price {
            if q - p > x {
                c++
                p = q
            }
        }
        return c < k
    })
}