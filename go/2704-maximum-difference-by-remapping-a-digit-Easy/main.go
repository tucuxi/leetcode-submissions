func minMaxDifference(num int) int {
    if num < 10 {
        return 9
    }
    s := []byte(strconv.Itoa(num))
    t := make([]byte, len(s))
    copy(t, s)
    i := 0
    for i < len(s) && s[i] == '9' {
        i++
    }
    if i < len(s) {
        replace(s, s[i], '9')
    }
    max, _ := strconv.Atoi(string(s))
    j := 0
    for j < len(t) && t[j] == '0' {
        j++
    }
    if j < len(t) {
        replace(t, t[j], '0')
    }
    min, _ := strconv.Atoi(string(t))
    return max - min
}

func replace(b []byte, v1, v2 byte) {
    for i := range b {
        if b[i] == v1 {
            b[i] = v2
        }
    }
}