func appendCharacters(s string, t string) int {
    j := 0
    for _, r := range s {
        if j == len(t) {
            return 0
        }
        if r == rune(t[j]) {
            j++
        }
    }
    return len(t) - j
}