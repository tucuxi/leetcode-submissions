func countMonobit(n int) int {
    res := 1
    for b := 1; b <= n; b = b << 1 | 1 {
        res++
    }
    return res
}