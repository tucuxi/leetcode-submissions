func findAndReplacePattern(words []string, pattern string) []string {
    res := []string{}
    
    outer:
    for _, w := range words {
        hp := make(map[byte]byte)
        hw := make(map[byte]byte)
        for i := range pattern {
            hpi, exists := hp[pattern[i]]
            if exists {
                if hpi != w[i] {
                    continue outer
                }
            } else {
                hp[pattern[i]] = w[i]
            }
            hwi, exists := hw[w[i]]
            if exists {
                if hwi != pattern[i] {
                    continue outer
                }
            } else {
                hw[w[i]] = pattern[i]
            }
        }
        res = append(res, w)
    }
    
    return res
}
