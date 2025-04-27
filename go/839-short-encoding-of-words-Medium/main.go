func minimumLengthEncoding(words []string) int {
    wm := make(map[string]struct{})
    for _, w := range words {
        wm[w] = struct{}{}
    }
    for w, _ := range wm {
        for k := 1; k < len(w); k++ {
            delete(wm, w[k:])
        }
    }
    res := 0
    for w, _ := range wm {
        res += len(w) + 1
    }
    return res
}