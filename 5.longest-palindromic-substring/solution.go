func longestPalindrome(s string) string {
    first, last := 0, 0
    for i := 1; i < len(s); i++ {
        if s[i - 1] == s[i] { 
            if a, b := expand(s, i - 1, i); b - a > last - first {
                first, last = a, b
            }
        }
        if  a, b := expand(s, i, i); b - a > last - first {
            first, last = a, b
        }
    }
    return s[first:last + 1]
}

func expand(s string, first int, last int) (int, int) {
    for first > 0 && last + 1 < len(s) && s[first - 1] == s[last + 1] {
        first--
        last++
    }
    return first, last
}