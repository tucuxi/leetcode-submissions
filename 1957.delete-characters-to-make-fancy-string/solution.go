func makeFancyString(s string) string {
    var sb strings.Builder
    for i := range s {
        if i > len(s) - 3 || s[i] != s[i+1] || s[i] != s[i+2] {
            sb.WriteByte(s[i])
        } 
    }
    return sb.String()
}