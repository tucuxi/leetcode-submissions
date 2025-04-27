func shortestPalindrome(s string) string {
    left := 0
    for right := len(s) - 1; right >= 0; right-- {
        if s[right] == s[left] {
            left++
        }
    }
    if left == len(s) {
        return s
    }
    nonPalindromicSuffix := s[left:]
    reverseSuffix := []byte(nonPalindromicSuffix)
    slices.Reverse(reverseSuffix)
    return string(reverseSuffix) + shortestPalindrome(s[:left]) + nonPalindromicSuffix
}