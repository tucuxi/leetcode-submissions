func removeTrailingZeros(num string) string {
    i := len(num)
    for i > 1 && num[i - 1] == '0' {
        i--
    }
    return num[:i]
}