func removeOuterParentheses(s string) string {
    var res strings.Builder
    k := 0
    for i := range s {
        if s[i] == '(' {
            k++
        }
        if k > 1 {
            res.WriteByte(s[i])
        }
        if s[i] == ')' {
            k--
        }
    }
    return res.String()
}