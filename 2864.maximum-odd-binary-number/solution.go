func maximumOddBinaryNumber(s string) string {
    b := []byte(s)
    last := len(b) - 1
    r := last
    for ; r > 0 && b[r] == '0'; r-- {
    }
    b[last], b[r] = b[r], b[last]
    r--
    for l := 0; l < r;  {
        if b[r] == '1' {
            b[r], b[l] = b[l], b[r]
            l++
        } else {
            r--
        }
    }
    return string(b)
}