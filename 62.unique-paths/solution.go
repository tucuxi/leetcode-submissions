func uniquePaths(m int, n int) int {
    return binom(m + n - 2, min(m, n) - 1)
}

func binom(n, k int) int {
    if k == 0 {
        return 1
    }
    return (n - k + 1) * binom(n, k - 1) / k
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}