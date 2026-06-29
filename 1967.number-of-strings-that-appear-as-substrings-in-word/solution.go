func numOfStrings(patterns []string, word string) int {
    res := 0
    for _, p := range patterns {
        if strings.Index(word, p) >= 0 {
            res++
        }
    }
    return res
}