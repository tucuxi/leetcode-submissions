func minimumFlips(n int) int {
    var b []bool

    for n > 0 {
        b = append(b, n & 1 == 1)
        n >>= 1
    }

    res := 0
    for i := range b {
        if b[i] != b[len(b) - i - 1] {
            res++
        }
    }

    return res
}