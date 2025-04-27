func grayCode(n int) []int {
    c := make([]int, 0, 1 << n)
    c = append(c, 0)
    for p := 1; p < 1 << n; p *= 2 {
        for i := p - 1; i >= 0; i-- {
            c = append(c, c[i] | p)
        }
    }
    return c
}