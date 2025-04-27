func getDescentPeriods(prices []int) int64 {
    previousPrice := 0
    previousStart := -1
    res := int64(0)

    for i, p := range prices {
        if p != previousPrice - 1 {
            previousStart = i
        }
        res += int64(i - previousStart + 1)
        previousPrice = p
    }

    return res
}