func isSubsequence(s string, t string) bool {
    for i, j := 0, 0; i < len(s); j++ {
        if j == len(t) {
            return false
        }
        if s[i] == t[j] {
            i++
        }
    }
    return true
}