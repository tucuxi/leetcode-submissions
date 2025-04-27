func findAllConcatenatedWordsInADict(words []string) []string {
    res := []string{}
    dict := map[string]bool{}
    for _, w := range words {
        dict[w] = true
    }
    for _, w := range words {
        dict[w] = false
        dp := make([]bool, len(w) + 1)
        dp[0] = true
        for i := 1; i <= len(w); i++ {
            for j := 0; j < i && !dp[i]; j++ {
                dp[i] = dp[j] && dict[w[j:i]]
            } 
        }
        if dp[len(w)] {
            res = append(res, w)
        }
        dict[w] = true
    }
    return res
}