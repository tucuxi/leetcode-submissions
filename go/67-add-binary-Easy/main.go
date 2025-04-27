func addBinary(a string, b string) string {
    r := make([]byte, 0, len(a) + len(b) + 1)
    var c byte
    for i, j := len(a) - 1, len(b) - 1; i >= 0 || j >= 0; {
        s := c
        if i >= 0 {
            s += a[i] - '0'
            i-- 
        }
        if j >= 0 {
            s += b[j] - '0'
            j--
        }
        c = (s & 2) >> 1
        s &= 1
        r = append(r, s + '0')
    }
    if c > 0 {
        r = append(r, '1')
    }
    for p, q := 0, len(r) - 1; p < q; p, q = p + 1, q - 1 {
        r[p], r[q] = r[q], r[p]
    }
    return string(r)
}