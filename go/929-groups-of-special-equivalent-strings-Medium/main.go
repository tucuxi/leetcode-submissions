func numSpecialEquivGroups(words []string) int {
    h := make(map[[2][26]int16]int)
    for _, w := range words {
        var v [2][26]int16
        for i := 0; i < len(w); i++ {
            v[i % 2][w[i] - 'a']++ 
        }
        h[v]++
    }
    return len(h)
}