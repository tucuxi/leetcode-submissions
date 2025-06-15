func removeDuplicates(s string) string {
    t := make([]byte, 0, len(s))
    tl := 0
    for i := range s {
        if tl > 0 && s[i] == t[tl-1] {
            tl--
            t = t[:tl]
        } else {
            tl++
            t = append(t, s[i])
        }
    }
    return string(t)
}