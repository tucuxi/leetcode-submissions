func numberOfChild(n int, k int) int {
    m := n - 1
    if k / m % 2 == 0 {
        return k % m
    }
    return m - k % m
}