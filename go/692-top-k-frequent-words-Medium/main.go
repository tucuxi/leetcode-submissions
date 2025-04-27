func topKFrequent(words []string, k int) []string {
    hist := map[string]int{}
    for _, w := range words {
        hist[w]++
    }
    uniq := []string{}
    for w := range hist {
        uniq = append(uniq, w)
    }
    sort.Slice(uniq, func(i, j int) bool {
        return hist[uniq[i]] == hist[uniq[j]] &&
            uniq[i] < uniq[j] ||
            hist[uniq[i]] > hist[uniq[j]] 
    })
    return uniq[:k]
}