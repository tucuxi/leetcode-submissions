func countBinarySubstrings(s string) int {
    p, q := 1, 0
    res := 0
    for i := 1; i < len(s); i++ {
        if s[i] == s[i - 1] {
            p++
        } else {
            p, q = 1, p
        }
        if q >= p {
            res++
        }
    }
    return res
}