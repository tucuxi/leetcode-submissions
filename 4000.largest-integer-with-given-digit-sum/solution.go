func largestInteger(n int, s int) int {
    res := 0
    for i := n; i > 0; i-- {
        d := min(s, 9)
        res = 10*res + d
        s -= d
    }
    if s > 0 {
        return -1
    }
    return res
}