func replaceDigits(s string) string {
    var sb strings.Builder
    for i := 0; i < len(s); i++ {
        if i % 2 == 0 {
            sb.WriteByte(s[i])
        } else {
            sb.WriteByte(shift(s[i-1], s[i] - '0'))
        }
    }
    return sb.String()
}

func shift(c, x byte) byte {
    return c + x
}