func minRemoveToMakeValid(s string) string {
    sb := make([]byte, 0, len(s))
    open := 0
    for i := range s {
        switch s[i] {
            case '(':
                open++
            case ')':
                if open == 0 {
                    continue
                }
                open--
        }
        sb = append(sb, s[i])
    }
    for i := len(sb) - 1; open > 0; i-- {
        if sb[i] == '(' {
            sb = append(sb[:i], sb[i+1:]...)
            open--
        }
    }
    return string(sb)
}