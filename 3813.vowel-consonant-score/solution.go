func vowelConsonantScore(s string) int {
    c, v := 0, 0
    for _, ch := range s {
        if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
            v++
        } else if ch >= 'a' && ch <= 'z' {
            c++
        }
    }
    if c > 0 {
        return v / c
    }
    return 0
}