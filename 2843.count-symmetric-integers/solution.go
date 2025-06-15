func countSymmetricIntegers(low int, high int) int {
    res := 0
    for i := max(low, 11); i <= min(high, 99); i++ {
        if i / 10 == i % 10 {
            res++
        }
    }
    for i := max(low, 1001); i <= high; i++ {
        h := i / 100
        sh := h / 10 + h % 10
        l := i % 100
        sl := l / 10 + l % 10
        if sh == sl {
            res++
        }
    }
    return res
}