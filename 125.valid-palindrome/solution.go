func isPalindrome(s string) bool {
    var b strings.Builder
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            b.WriteRune(unicode.ToLower(r))
        }
    }
    t := b.String()
    i, j := 0, len(t) - 1
    for i < j {
        if t[i] != t[j] {
            return false
        }
        i++
        j--
    }
    return true
}