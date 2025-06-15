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
                return reversedIfTrue(b.String(), inside), pos + 1
            default:
                b.WriteByte(s[pos])
                pos++
        }
    }

    return reversedIfTrue(b.String(), inside), pos
}

func reversedIfTrue(s string, reverse bool) string {
    if !reverse {
        return s
    }

    var b strings.Builder

    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    } 
    return b.String()
}