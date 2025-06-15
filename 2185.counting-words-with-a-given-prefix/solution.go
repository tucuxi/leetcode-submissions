func prefixCount(words []string, pref string) int {
    res := 0
    for _, w := range words {
        if strings.HasPrefix(w, pref) {
            res++
        }
    }
    return res
}