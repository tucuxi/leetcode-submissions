func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    run := 0
    var sb strings.Builder
    var ch byte
    for s := countAndSay(n - 1); len(s) > 0; s = s[1:] {
        if s[0] != ch {
            writeEncoding(&sb, ch, run)
            run = 0
        }
        ch = s[0]
        run++
    }
    writeEncoding(&sb, ch, run)
    return sb.String()
}


func writeEncoding(b *strings.Builder, ch byte, length int) {
    b.WriteString(encode(ch, length))
}

func encode(ch byte, length int) string {
    if length == 0 {
        return ""
    }
    return fmt.Sprintf("%d%c", length, ch)
}