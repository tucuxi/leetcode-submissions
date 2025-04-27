func removeOccurrences(s string, part string) string {
    st := make([]byte, 0, len(s))
    
    outer:
    for i := range s {
        st = append(st, s[i])
        if len(st) >= len(part) {
            k := len(st) - len(part)
            for j := 0; j < len(part); j++ {
                if st[k + j] != part[j] {
                    continue outer
                }
            }
            st = st[:k]
        }
    }
    return string(st)
}