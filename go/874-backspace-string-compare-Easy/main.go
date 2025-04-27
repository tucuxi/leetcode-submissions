func backspaceCompare(s string, t string) bool {
    return norm(s) == norm(t)
}
 
func norm(s string) string {
    var (
        b strings.Builder
        skip int
    )
    for i := len(s) - 1; i >= 0; i-- {
        switch {
            case s[i] == '#': skip++
            case skip > 0: skip--
            default: b.WriteByte(s[i])
        }
    }
    return b.String()
}