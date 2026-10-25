func reverseParentheses(s string) string {
    res, _ := dfs(s, 0, false)
    return res
}

func dfs(s string, pos int, inside bool) (string, int) {
    var b strings.Builder

    for pos < len(s) {
        switch s[pos] {
            case '(':
                t, i := dfs(s, pos + 1, true)
                b.WriteString(t)
                pos = i
            case ')':
                return toString(b, inside), pos + 1
            default:
                b.WriteByte(s[pos])
                pos++
        }
    }

    return toString(b, inside), pos
}

func toString(b strings.Builder, reverse bool) string {
    s := b.String()

    if !reverse {
        return s
    }

    r := make([]byte, len(s))

    for i := range r {
        r[i] = s[len(s) - 1 - i]
    }
    
    return string(r)
}