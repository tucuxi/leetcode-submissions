func largestAltitude(gain []int) int {
    a := 0
    m := 0
    for _, g := range gain {
        a += g
        m = max(a, m)
    }
    return m
}