func evaluate(s string, knowledge [][]string) string {
    kv := make(map[string]string)
    for _, p := range knowledge {
        kv[p[0]] = p[1]
    }

    var res strings.Builder
    k := 0

    for i := range s {
        switch s[i] {
        case '(':
            k = i+1
        case ')':
            if v, ok:= kv[s[k:i]]; ok {
                res.WriteString(v)
            } else {
                res.WriteString("?")
            }
            k = 0
        default:
            if k == 0 {
                res.WriteByte(s[i])
            }
        }
    }
    return res.String()
}