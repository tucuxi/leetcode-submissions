func parseBoolExpr(expression string) bool {
    v, _ := expr(expression)
    return v
}

func notExpr(s string) (bool, string) {
    v, s := expr(s)
    return !v, s[1:]
}

func andExpr(s string) (bool, string) {
    v, s := expr(s)
    for s[0] == ',' {
        w, t := expr(s[1:])
        v = v && w
        s = t
    }
    return v, s[1:]
}

func orExpr(s string) (bool, string) {
    v, s := expr(s)
    for s[0] == ',' {
        w, t := expr(s[1:])
        v = v || w
        s = t
    }
    return v, s[1:]
}

func expr(s string) (bool, string) {
    switch s[0] {
        case 't': return true, s[1:]
        case 'f': return false, s[1:]
        case '!': return notExpr(s[2:])
        case '&': return andExpr(s[2:])
        case '|': return orExpr(s[2:])
        default: panic("Invalid input")
    }
}