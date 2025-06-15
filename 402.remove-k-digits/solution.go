func removeKdigits(num string, k int) string {
    var st []byte
    for i := range num {
        for k > 0 && len(st) > 0 && st[len(st) - 1] > num[i] {
            st = st[:len(st) - 1]
            k--
        }
        st = append(st, num[i])
    }
    if k >= len(st) {
        return "0"
    }
    st = st[:len(st) - k]
    for len(st) > 1 && st[0] == '0' {
        st = st[1:]
    }
    return string(st)
}