func greatestLetter(s string) string {
    uc := [26]bool{}
    lc := [26]bool{}
    for _, c := range s {
        if c >= 'A' && c <= 'Z' {
            uc[c - 'A'] = true
        } else if c >= 'a' && c <= 'z' {
            lc[c - 'a'] = true
        }
    }
    for i := len(uc) - 1; i >= 0; i-- {
        if uc[i] && lc[i] {
            return string(i + 'A')
        }
    }
    return ""
}