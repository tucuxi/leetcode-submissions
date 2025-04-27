func hasAlternatingBits(n int) bool {
    for p := n; p > 0; {
        q := p >> 1
        if (p & 1 == q & 1) {
            return false
        }
        p = q
    }
    return true
}