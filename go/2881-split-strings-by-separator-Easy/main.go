func splitWordsBySeparator(words []string, separator byte) []string {
    var res []string
    for _, w := range words {
        var b strings.Builder
        for _, r := range w {
            c := byte(r)
            if c == separator {
                if b.Len() > 0 {
                    res = append(res, b.String())
                    b.Reset()
                }
            } else {
                b.WriteByte(c)
            }
        }
        if b.Len() > 0 {
            res = append(res, b.String())
        }
    }
    return res
}