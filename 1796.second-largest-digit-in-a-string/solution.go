func secondHighest(s string) int {
    digits := [10]bool{}
    for _, c := range s {
        if c >= '0' && c <= '9' {
            digits[c - '0'] = true
        }
    }
    cnt := 0
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] {
            cnt++
            if cnt == 2 {
                return i
            }
        }
    }
    return -1
}