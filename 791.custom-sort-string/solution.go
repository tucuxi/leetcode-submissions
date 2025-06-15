func customSortString(order string, s string) string {
    index := make(map[byte]int)
    for i := range order {
        index[order[i]] = i
    }
    res := []byte(s)
    slices.SortFunc(res, func(a, b byte) int {
        return index[a] - index[b]
    })
    return string(res)
}