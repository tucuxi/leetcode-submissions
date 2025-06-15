func reformatNumber(number string) string {
    var b strings.Builder
    for _, c := range number {
        if c >= '0' && c <= '9' {
            b.WriteByte(byte(c))
        }
    }
    s := b.String()
    b.Reset()
    i := 0
    for i < len(s) {
        switch len(s) - i {
        case 2:
            b.WriteString(s[i:i+2])
            i += 2
        case 3:
            b.WriteString(s[i:i+3])
            i += 3
        case 4:
            b.WriteString(s[i:i+2])
            b.WriteByte('-')
            b.WriteString(s[i+2:i+4])
            i += 4
        default:
            b.WriteString(s[i:i+3])
            b.WriteByte('-')
            i += 3
        }
    }
    return b.String()
}