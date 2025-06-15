func longestWord(words []string) string {
    h := make(map[string]bool)
    for _, w := range words {
        h[w] = true
    }
    
    res := ""
    
    outer:
    for _, w := range words {
        if len(w) < len(res) || len(w) == len(res) && w > res {
            continue
        }
        for i := 1; i < len(w); i++ {
            if !h[w[:i]] {
                continue outer
            }
        }
        res = w
    }
    return res
}