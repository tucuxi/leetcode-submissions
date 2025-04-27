func getEncryptedString(s string, k int) string {
    var sb strings.Builder

    for i := range s {
        j := (i + k) % len(s)
        sb.WriteByte(s[j])
    }
    return sb.String()
}