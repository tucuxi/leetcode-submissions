func minimumPartition(s string, k int) int {
    res := 1
    n := 0
    for i := range s {
        d := int(s[i] - '0')
        if d > k {
            return -1
        }
        n = 10 * n + d
        if n > k {
            res++
            n = d
        }
    }
    return res
}