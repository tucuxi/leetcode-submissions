func getSum(a int, b int) int {
    a, b = a ^ b, a & b
    for b != 0 {
        h := b << 1
        a, b = a ^ h, a & h
    }
    return a
}