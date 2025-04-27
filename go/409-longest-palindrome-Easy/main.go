func longestPalindrome(s string) int {
    h := make(map[rune]int)
    for _, r := range s {
        h[r]++
    }
    n := 1
    for _, v := range h {
        n += v & ^1
    }
    if n > len(s) {
        return len(s)
    }
    return n
}