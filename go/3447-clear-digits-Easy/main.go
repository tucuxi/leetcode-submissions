func clearDigits(s string) string {
    b := make([]byte, 0, len(s))
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            b = b[:len(b) - 1]
        } else {
            b = append(b, s[i])
        }
    }
    return string(b)
}