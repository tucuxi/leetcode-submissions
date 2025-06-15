func makeGood(s string) string {
    t := make([]byte, 0, len(s))
    t = append(t, s[0])
    for p, q := 1, 0; p < len(s); p++ {
        if q >= 0 && abs(int(t[q]) - int(s[p])) == 32 {
            t = t[:q]
            q--
        } else {
            t = append(t, s[p])
            q++
        }
    }
    return string(t)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}