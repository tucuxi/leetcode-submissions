func strongPasswordCheckerII(password string) bool {
    if (len(password) < 8) {
        return false
    }
    lc, uc, dg, sp := false, false, false, false
    for _, c := range password {
        switch {
            case c >= 'a' && c <= 'z': lc = true
            case c >= 'A' && c <= 'Z': uc = true
            case c >= '0' && c <= '9': dg = true
            default: sp = true
        }
    }
    if !lc || !uc || !dg || !sp {
        return false
    }
    for i := 1; i < len(password); i++ {
        if password[i - 1] == password[i] {
            return false
        }
    }
    return true
}