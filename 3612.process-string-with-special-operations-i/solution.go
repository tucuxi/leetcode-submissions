func processStr(s string) string {
    var b []byte

    for i := range s {
        switch s[i] {
        case '#':
            duplicate(&b)
        case '%':
            reverse(b)
        case '*':
            removeLast(&b)
        default:
            b = append(b, s[i])
        }
    }

    return string(b)
}

func duplicate(b *[]byte) {
    *b = append(*b, *b...)
}

func reverse(b []byte) {
    i := 0
    j := len(b) - 1
    for i < j {
        b[i], b[j] = b[j], b[i]
        i++
        j--
    }
}

func removeLast(b *[]byte) {
    if len(*b) > 0 {
        *b = (*b)[:len(*b)-1]
    }
}