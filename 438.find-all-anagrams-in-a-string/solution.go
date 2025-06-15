func findAnagrams(s string, p string) []int {
    res := []int{}
    if len(s) < len(p) {
        return res
    }
    h := [26]int{}
    eval := func(i int) {
        for _, count := range h {
            if count != 0 {
                return
            }
        }
        res = append(res, i)
    }
    for i := range p {
        h[p[i] - 'a']--
        h[s[i] - 'a']++
    }
    eval(0)
    for i := len(p); i < len(s); i++ {
        h[s[i - len(p)] - 'a']--
        h[s[i] - 'a']++
        eval(i - len(p) + 1)
    }
    return res
}