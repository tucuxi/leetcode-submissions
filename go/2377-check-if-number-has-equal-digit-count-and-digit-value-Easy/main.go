func digitCount(num string) bool {
    digits := [10]byte{}
    for i := range num {
        digits[num[i] - '0']++
    }
    for i := range num {
        if num[i] - '0' != digits[i] {
            return false
        }
    }
    return true
}