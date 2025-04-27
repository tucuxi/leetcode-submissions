func countKeyChanges(s string) int {
    res := 0
    key := s[0] | 0x20
    for i := 1; i < len(s); i++ {
        cur := s[i] | 0x20
        if cur != key {
            res++
            key = cur
        }
    }
    return res
}