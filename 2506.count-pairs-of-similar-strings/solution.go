func similarPairs(words []string) int {
    m := map[[26]bool]int{}
    for _, word := range words {
        key := [26]bool{}
        for i := range word {
            key[word[i] - 'a'] = true 
        }
        m[key]++
    }
    res := 0
    for _, count := range m {
        res += count * (count - 1) / 2
    }
    return res
}