func longestNiceSubstring(s string) string {
    if len(s) < 2 {
        return ""
    }
    letters := map[byte]bool{}
    for i := range s {
        letters[s[i]] = true
    }
    for i := range s {
        swapcase := s[i] ^ 0x20
        if !letters[swapcase] {
            s1 := longestNiceSubstring(s[:i])
            s2 := longestNiceSubstring(s[i+1:])
            if len(s1) < len(s2) {
                return s2
            }
            return s1
        }
    }
    return s
}