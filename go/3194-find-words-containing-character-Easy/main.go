func findWordsContaining(words []string, x byte) []int {
    var res []int

    for i, w := range words {
        if strings.IndexByte(w, x) >= 0 {
            res = append(res, i)
        }
    }
    return res
}