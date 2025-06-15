func rearrangeCharacters(s string, target string) int {
    h := map[byte]int{}
    for i := range s {
        h[s[i]]++
    }
    res := 0
    for {
        for i := range target {
            if h[target[i]] == 0 {
                return res
            }
            h[target[i]]--
        }
        res++
    }
}