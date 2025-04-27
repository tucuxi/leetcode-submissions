func stringHash(s string, k int) string {
    var b strings.Builder
    for i := 0; i < len(s); i += k {
        var h byte
        for j := range k {
            h = (h + s[i + j] - 'a') % 26
        }
        b.WriteByte(h + 'a')
    }
    return b.String()
}