func countTriples(n int) int {
    res := 0
    for a := 1; a <= n; a++ {
        for b := 1; b <= n; b++ {
            target := a * a + b * b
            c := sort.Search(n + 1, func(c int) bool {
                return c * c >= target
            })
            if c <= n && c * c == target {
                res++
            }
        }
    }
    return res
}