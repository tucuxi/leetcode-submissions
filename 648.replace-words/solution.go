func replaceWords(dictionary []string, sentence string) string {
    h := make(map[string]bool)
    for _, d := range dictionary {
        h[d] = true
    }
    words := strings.Split(sentence, " ")
    var res []string
    for _, w := range words {
        i := 1
        for i < len(w) && !h[w[:i]] {
            i++
        }
        res = append(res, w[:i])
    }
    return strings.Join(res, " ")
}