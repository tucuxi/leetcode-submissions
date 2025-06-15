func oddString(words []string) string {
    res := ""
    for i := 1; i < len(words); i++ {
        if !eq(words[i], words[0]) {
            if res == "" {
                res = words[i]
            } else {
                return words[0]
            }
        }
    }
    return res
}

func eq(s1, s2 string) bool {
    for i := 1; i < len(s1); i++ {
        if s1[i] - s1[i - 1] != s2[i] - s2[i - 1] {
            return false
        }
    }
    return true
}