func checkOnesSegment(s string) bool {
    i := 0
    for s[i] == '1' {
        i++
        if i == len(s) {
            return true
        }
    }
    for s[i] == '0' {
        i++
        if i == len(s) {
            return true
        }
    }
    return false
}