func checkAlmostEquivalent(word1 string, word2 string) bool {
    h := map[byte]int{}
    for i := range word1 {
        h[word1[i]]++
        h[word2[i]]--
    }
    for _, f := range h {
        if f > 3 || f < -3 {
            return false
        }
    }
    return true
}
