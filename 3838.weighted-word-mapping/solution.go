func mapWordWeights(words []string, weights []int) string {
    res := make([]byte, 0, len(words))

    for _, word := range words {
        weight := 0
        for i := range word {
            weight += weights[word[i] - 'a']
        }
        ch := 'z' - byte(weight % 26)
        res = append(res, ch)
    }

    return string(res)
}