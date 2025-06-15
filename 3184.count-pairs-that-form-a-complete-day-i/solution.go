func countCompleteDayPairs(hours []int) int {
    var count [24]int
    res := 0

    for _, h := range hours {
        h %= 24
        res += count[(24-h) % 24]
        count[h]++
    }

    return res
}