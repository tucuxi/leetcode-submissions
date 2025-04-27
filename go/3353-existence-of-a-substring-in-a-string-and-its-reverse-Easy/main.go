func isSubstringPresent(s string) bool {
    di := make(map[[2]byte]bool)
    for i := 1; i < len(s); i++ {
        f := [...]byte{ s[i-1], s[i] }
        di[f] = true
        b := [...]byte{ s[i], s[i-1] }
        if di[b] {
            return true
        }
    }
    return false
}