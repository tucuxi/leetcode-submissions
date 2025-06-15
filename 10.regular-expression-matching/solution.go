func isMatch(s string, p string) bool {
    if p == "" {
        return s == ""
    }
    if len(p) > 1 && p[1] == '*' {
        return matchStar(p[0], s, p[2:])
    }
    if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
        return isMatch(s[1:], p[1:])
    }
    return false
}

func matchStar(c byte, s string, p string) bool {
    for {
        if isMatch(s, p) {
            return true
        }
        if s == "" || s[0] != c && c != '.' {
            return false
        }
        s = s[1:]
    }
}