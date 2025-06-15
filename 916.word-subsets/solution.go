func wordSubsets(words1 []string, words2 []string) []string {
    var h2 [26]int
    for _, w2 := range words2 {
        var w2c [26]int
        for _, c := range w2 {
            w2c[c - 'a']++
        }
        for i := range w2c {
            h2[i] = max(h2[i], w2c[i])
        }
    }

    var res []string

    outer:
    for _, w1 := range words1 {
        var w1c [26]int
        for _, c := range w1 {
            w1c[c- 'a']++
        }
        for i := range w1c {
            if w1c[i] < h2[i] {
                continue outer
            } 
        }
        res = append(res, w1)
    }

    return res
}