func uncommonFromSentences(s1 string, s2 string) []string {
    h := make(map[string]int)
    for _, w := range strings.Fields(s1) {
        h[w]++
    }
    for _, w := range strings.Fields(s2) {
        h[w]++
    }
    var res []string
    for w, n := range h {
        if n == 1 {
            res = append(res, w)
        }
    }
    return res
}