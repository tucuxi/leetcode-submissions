func minDeletions(s string) int {
    h := make(map[rune]int)
    for _, r := range s {
        h[r]++
    }
    f := make(map[int]int)
    for _, n := range h {
        f[n]++
    }
    res := 0
    for r := 'a'; r <= 'z'; r++ {
        k := h[r]
        for k > 0 && f[k] > 1 {
            res++
            f[k]--
            k--
            f[k]++
        }
    }
    return res
}