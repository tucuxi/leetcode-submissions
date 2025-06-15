func divideString(s string, k int, fill byte) []string {
    res := []string{}
    i := 0
    for i <= len(s) - k {
        res = append(res, s[i:i+k])
        i += k
    }
    if i < len(s) {
        var sb strings.Builder
        sb.WriteString(s[i:])
        for sb.Len() < k {
            sb.WriteByte(fill)
        }
        res = append(res, sb.String())
    } 
    return res
}