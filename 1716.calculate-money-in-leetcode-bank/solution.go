func totalMoney(n int) int {
    s := 0
    for i := 0; i < n; i++ {
        s += i/7 + i%7 + 1
    }
    return s
}