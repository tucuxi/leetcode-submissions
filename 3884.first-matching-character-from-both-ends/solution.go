func firstMatchingIndex(s string) int {
    i, j := 0, len(s) - 1
    for i <= j {
        if s[i] == s[j] {
            return i
        }
        i++
        j--
    }
    return -1
}