func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
    res := 0
    prev := math.MinInt
    for _, p := range pairs {
        if prev < p[0] {
            res++
            prev = p[1]
        }
    }
    return res
}   