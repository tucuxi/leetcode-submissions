func smallestSubsequence(s string) string {
    f := [26]int{}
    for i := range s {
        f[s[i] - 'a']++
    }
    res := make([]byte, 0, 26)
    m := 0
    for i := range s {
        c := s[i] - 'a'
        f[c]--
        if m & (1 << c) == 0 {
            for r := len(res) - 1; r >= 0 && s[i] < res[r] && f[res[r] - 'a'] > 0; r-- {
                m ^= 1 << (res[r] - 'a')
                res = res[:r]
            }
            m |= 1 << c
            res = append(res, s[i])
        }
    }
    return string(res)    
}