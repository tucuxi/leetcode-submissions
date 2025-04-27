func findLatestTime(s string) string {
    b := []byte(s)
    if b[0] == '?' {
        if b[1] >= '2' && b[1] <= '9' {
            b[0] = '0'
        } else {
            b[0] = '1'
        }
    }
    if b[1] == '?' {
        if b[0] == '0' {
            b[1] = '9'
        } else {
            b[1] = '1'
        }
    }
    if b[3] == '?' {
        b[3] = '5'
    }
    if b[4] == '?' {
        b[4] = '9'
    }
    return string(b)
}