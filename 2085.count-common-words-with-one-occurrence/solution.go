func countWords(words1 []string, words2 []string) int {
    h1 := hist(words1)
    h2 := hist(words2)
    res := 0
    for word, cnt := range h1 {
        if cnt == 1 {
            if h2[word] == 1 {
                res++
            }
        }
    }
    return res
}

func hist(words []string) map[string]int {
    h := map[string]int{}
    for _, w := range words {
        h[w]++
    }
    return h
}