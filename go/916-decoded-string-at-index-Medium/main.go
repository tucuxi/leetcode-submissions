func decodeAtIndex(s string, k int) string {
    var (
        length int64
        i int
    )
    
    for length < int64(k) {
        if s[i] >= '0' && s[i] <= '9' {
            length *= int64(s[i] - '0')
        } else {
            length++
        }
        i++
    }
    
    for i > 0 {
        i--
        if s[i] >= '0' && s[i] <= '9' {
            length /= int64(s[i] - '0')
            k %= int(length)
        } else {
            if k == 0 || k == int(length) {
                return string(s[i])
            }
            length--
        }
    }
    
    panic("no solution")
}