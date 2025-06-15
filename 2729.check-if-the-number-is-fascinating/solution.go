func isFascinating(n int) bool {
    s := fmt.Sprintf("%d%d%d", n, 2 * n, 3 * n)
    digits := make([]int, 10)
    for _, c := range s {
        digits[c - '0']++
    }
    for d := 1; d < 10; d++ {
        if digits[d] != 1 {
            return false
        }
    }
    return true
}