func thousandSeparator(n int) string {
    if n == 0 {
        return "0"
    }
    var r strings.Builder
    for i := 0; n > 0; n /= 10 {
        if i > 0 && i % 3 == 0 {
            r.WriteByte('.')
        }
        r.WriteByte('0' + byte(n % 10))
        i++
    }
    return reverse(r.String())
}

func reverse(s string) string {
    var sb strings.Builder
    for i := len(s) - 1; i >= 0; i-- {
        sb.WriteByte(s[i])
    }
    return sb.String()
}