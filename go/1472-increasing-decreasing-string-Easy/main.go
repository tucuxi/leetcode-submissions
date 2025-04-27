func sortString(s string) string {
    var res strings.Builder
    chars := [26]int{}
    for i := range s {
        chars[s[i] - 'a']++
    }
    for res.Len() < len(s) {
        for i := 0; i < len(chars); i++ {
            if chars[i] > 0 {
                res.WriteByte('a' + byte(i))
                chars[i]--
            }
        }
        for i := len(chars) - 1; i >= 0; i-- {
            if chars[i] > 0 {
                res.WriteByte('a' + byte(i))
                chars[i]--
            }
        }
    }
    return res.String()
}
