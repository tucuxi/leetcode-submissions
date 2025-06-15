func concatenatedBinary(n int) int {
    res := 0
    l := 0
    for i := 1; i <= n; i++ {
        if i & (i - 1) == 0 {
            l++
        }
        res = ((res << l) + i) % 1_000_000_007
    }
    return res
}