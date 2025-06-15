func minEnd(n int, x int) int64 {
    res := int64(x)
    j := 0
    for i := 1; i <= n; i *= 2 {
        for x & (1 << j) != 0 {
            j++
        }
        if (n - 1) & i != 0 {
            res |= int64(1 << j)
        }
        j++
    }
    return res
}