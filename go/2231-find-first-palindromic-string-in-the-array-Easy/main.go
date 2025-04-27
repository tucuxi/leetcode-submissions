func firstPalindrome(words []string) string {
    for _, w := range words {
        if isPalindromic(w) {
            return w
        }
    }
    return ""
}

func isPalindromic(s string) bool {
    l := len(s) - 1
    for i := 0; i < len(s) / 2; i++ {
        if s[i] != s[l - i] {
            return false
        }
    }
    return true
}