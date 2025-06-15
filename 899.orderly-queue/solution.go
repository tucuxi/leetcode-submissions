func orderlyQueue(s string, k int) string {
    b := []byte(s)
    if k > 1 {
        sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
        return string(b)
    }
    m := make([]byte, len(b))
    copy(m, b)
    for _ = range s {
        if bytes.Compare(b, m) < 0 {
            copy(m, b)
        }
        b = append(b[1:], b[0])
    }
    return string(m)
}