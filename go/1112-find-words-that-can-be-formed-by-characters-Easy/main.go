func countCharacters(words []string, chars string) int {
    f := make([]int, 26)
    for _, c := range chars {
        f[c - 'a']++
    }
    
    h := make([]int, 26)
    res := 0
    
    outer:
    for _, w := range words {
        copy(h, f)
        for _, c := range w {
            k := c - 'a'
            if h[k] == 0 {
                continue outer
            }
            h[k]--
        }
        res += len(w)
    }
    
    return res
}