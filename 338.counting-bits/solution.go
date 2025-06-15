func countBits(n int) []int {
    res := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        res[i] = res[i / 2] + i % 2
    }
    return res
}