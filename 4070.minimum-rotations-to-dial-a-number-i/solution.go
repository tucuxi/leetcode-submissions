func minRotations(s string) int {
    res := 0
    p := 0
    for _, r := range s {
        q := int(r - '0')
        c := (p-q+10) % 10
        res += min(c, 10-c)
        p = q
    }
    return res
}