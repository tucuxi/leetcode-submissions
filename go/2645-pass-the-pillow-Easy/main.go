func passThePillow(n int, time int) int {
    p := 1 + time % (2 * (n - 1))
    if p > n {
        return 2 * n - p
    }
    return p
}