func decodeString(s string) string {
    var b strings.Builder
    p := 0

    var decodeBlock func() string
    decodeBlock = func() string {
        k := int(s[p] - '0')
        for p++; s[p] != '['; p++ {
            k = 10 * k + int(s[p] - '0')
        }
        var b strings.Builder
        for p++; s[p] != ']'; p++ {
            if s[p] >= '0' && s[p] <= '9' {
                b.WriteString(decodeBlock())
            } else {
                b.WriteByte(s[p])
            }
        }
        return strings.Repeat(b.String(), k)
    }

    for ; p < len(s); p++ {
        if s[p] >= '0' && s[p] <= '9' {
            b.WriteString(decodeBlock())
        } else {
            b.WriteByte(s[p])
        }
    }
   return b.String()
}
