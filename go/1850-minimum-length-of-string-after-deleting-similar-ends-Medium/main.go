func minimumLength(s string) int {
    l := 0
    r := len(s) - 1

    for l < r && s[l] == s[r] {
        c := s[l]
        for l <= r && s[l] == c {
            l++
        }
        for l <= r && s[r] == c {
            r--
        }
    }
    return r-l+1
}