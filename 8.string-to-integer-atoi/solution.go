func myAtoi(s string) int {
    const (
        minInt = -0x80000000
        maxInt = 0x7fffffff
    )
    i := 0
    for ; i < len(s) && s[i] == ' '; i++ {
    }
    var sign int64
    if i < len(s) {
        switch s[i] {
            case '+':
                sign = 1
                i++
            case '-':
                sign = -1
                i++
            default:
                sign = 1
        }
    }
    var n int64
    for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
        switch {
            case n == 0:
                n = sign * (int64(s[i]) - int64('0'))
            case n > 0:
                n = n * 10 + int64(s[i]) - int64('0')
                if n > maxInt {
                    n = maxInt
                    break
                }
            case n < 0:
                n = n * 10 - (int64(s[i]) - int64('0'))
                if n < minInt {
                    n = minInt
                    break
                }
        }
    }
    return int(n)
}