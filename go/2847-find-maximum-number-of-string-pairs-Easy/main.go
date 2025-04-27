func maximumNumberOfStringPairs(words []string) int {
    res := 0
    h := make(map[string]bool)
    for _, w := range words {
        if r := fmt.Sprintf("%c%c", w[1], w[0]); h[r] {
            res++
        } else {
            h[w] = true
        }
    }
    return res
}