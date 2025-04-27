func hasSameDigits(s string) bool {
    b := make([]byte, len(s))
    for i := range s {
        b[i] = s[i] - '0'
    }
    for len(b) > 2 {
        var c []byte
        for i := 1; i < len(b); i++ {
            c = append(c, (b[i - 1] + b[i]) % 10)
        }
        b = c
    }
    return b[0] == b[1]
}