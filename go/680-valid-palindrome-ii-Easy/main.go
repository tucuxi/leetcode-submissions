func validPalindrome(s string) bool {
    a, b := scan(s, 0, len(s)-1)
    if a >= b {
        return true
    }
    if c, d := scan(s, a+1, b); c >= d {
        return true
    }
    c, d := scan(s, a, b-1)
    return c >= d
}

func scan(s string, i, j int) (int, int) {
    for i < j && s[i] == s[j] {
        i++
        j--
    }
    return i, j
}