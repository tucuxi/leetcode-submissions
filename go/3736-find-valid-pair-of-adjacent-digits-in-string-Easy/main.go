func findValidPair(s string) string {
    var h [10]int
    for _, r := range s {
        h[int(r - '0')]++
    }
    for i := 1; i < len(s); i++ {
        if s[i] == s[i - 1] {
            continue
        }
        if a := int(s[i - 1] - '0'); h[a] != a {
            continue
        }
        if b := int(s[i] - '0'); h[b] != b {
            continue
        }
        return s[i - 1: i + 1]
    }
    return ""
}