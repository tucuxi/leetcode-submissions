func reformat(s string) string {
    var letters, digits strings.Builder
    for _, c := range s {
        if c >= '0' && c <= '9' {
            digits.WriteByte(byte(c))
        } else {
            letters.WriteByte(byte(c))
        }
    }
    ls, ds := letters.String(), digits.String()
    if abs(len(ls) - len(ds)) > 1 {
        return ""
    }
    var res strings.Builder
    if len(ls) <= len(ds) {
        for i := range ls {
            res.WriteByte(ds[i])
            res.WriteByte(ls[i])
        }
        if len(ls) < len(ds) {
            res.WriteByte(ds[len(ls)])
        }
    } else {
        for i := range ds {
            res.WriteByte(ls[i])
            res.WriteByte(ds[i])
        }
        res.WriteByte(ls[len(ds)])
    }
    return res.String()
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}