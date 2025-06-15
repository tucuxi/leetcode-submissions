func minChanges(s string) int {
    res := 0
    k := 0
    for i := range s {
        if s[i] != s[k] {
            r := (i - k) % 2
            res += r
            k = i + r
        }
    }
    return res
}