func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }
    outer:
    for i := range s {
        for j := range goal {
            if s[j] != goal[(i + j) % len(goal)] {
                continue outer
            }
        }
        return true
    }
    return false
}