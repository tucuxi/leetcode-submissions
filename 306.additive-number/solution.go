func isAdditiveNumber(num string) bool {
    for j := 1; j < min(18, len(num)); j++ {
        if j > 1 && num[0] == '0' {
            break
        }
        a, err := strconv.ParseUint(num[:j], 10, 64)
        if err != nil {
            panic(err)
        }
        for k := j+1; k < min(j+18, len(num)); k++ {
            if k-j > 1 && num[j] == '0' {
                break
            }
            b, err := strconv.ParseUint(num[j:k], 10, 64)
            if err != nil {
                panic(err)
            }
            if f(a, b, num[k:]) {
                return true
            }
        }
    }
    return false
}

func f(a, b uint64, s string) bool {
    c := strconv.FormatUint(a+b, 10)
    if len(c) > len(s) || c != s[:len(c)] {
        return false
    }
    r := s[len(c):]
    return len(r) == 0 || f(b, a+b, s[len(c):])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}