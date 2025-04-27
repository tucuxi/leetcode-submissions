func nthUglyNumber(n int) int {
    h := make([]int, n)
    h[0] = 1
    index2 := 0
    index3 := 0
    index5 := 0
    for i := 1; i < n; i++ {
        candidate2 := h[index2] * 2
        candidate3 := h[index3] * 3
        candidate5 := h[index5] * 5
        h[i] = min(candidate2, candidate3, candidate5)
        if h[i] == candidate2 {
            index2++
        }
        if h[i] == candidate3 {
            index3++
        }
        if h[i] == candidate5 {
            index5++
        }
    }
    return h[n-1]
}