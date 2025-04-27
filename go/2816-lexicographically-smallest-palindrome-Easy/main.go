func makeSmallestPalindrome(s string) string {
    p := []byte(s)
    for i, j := 0, len(p) - 1; i < j; i, j = i + 1, j - 1 {
        if p[i] > p[j] {
            p[i] = p[j]
        } else {
            p[j] = p[i]
        }
    }
    return string(p)
}