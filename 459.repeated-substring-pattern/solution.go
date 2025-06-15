func repeatedSubstringPattern(s string) bool {
    for l := 1; l <= len(s) / 2; l++ {
        if len(s) % l == 0 {
            match := true
            for i := l; match && i < len(s); i += l {
                match = s[:l] == s[i:i+l]
            }
            if match {
                return true
            }
        }
    }
    return false
}