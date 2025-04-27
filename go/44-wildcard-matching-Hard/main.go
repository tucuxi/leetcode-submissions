func isMatch(s string, p string) bool {
    i := 0
    j := 0
    star := -1
    match := 0

    for i < len(s) {
        switch {
        case j < len(p) && (p[j] == s[i] || p[j] == '?'):
            i++
            j++
        case j < len(p) && p[j] == '*':
            star = j
            match = i
            j++
        case star >= 0:
            match++
            i = match
            j = star + 1
        default:
            return false
        }
    }
    for j < len(p) && p[j] == '*' {
        j++
    }

    return j == len(p)
}