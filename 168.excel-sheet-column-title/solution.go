const letters = 26

func convertToTitle(columnNumber int) string {
    var sb strings.Builder
    c := columnNumber
    for c > 0 {
        c--
        sb.WriteByte('A' + byte(c % letters))
        c /= letters
    }
    r := []byte(sb.String())
    for i, j := 0, len(r) - 1; i < j; i, j = i + 1, j - 1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}