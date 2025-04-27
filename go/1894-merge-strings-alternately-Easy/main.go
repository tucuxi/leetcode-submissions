func mergeAlternately(word1 string, word2 string) string {
    var b strings.Builder
    for len(word1) + len(word2) > 0 {
        if len(word1) > 0 {
            b.WriteByte(word1[0])
            word1 = word1[1:]
        }
        if len(word2) > 0 {
            b.WriteByte(word2[0])
            word2 = word2[1:]
        }
    }
    return b.String()
}