func alternateDigitSum(n int) int {
    res := 0
    sign := 1
    for n > 0 {
        res += (n % 10) * sign
        sign = -sign
        n /= 10
    }
    return -sign * res
}