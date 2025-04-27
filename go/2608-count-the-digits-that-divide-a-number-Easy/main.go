func countDigits(num int) int {
    res := 0
    for n := num; n > 0; n /= 10 {
        digit := n % 10
        if digit > 0 && num % digit == 0 {
            res++
        }
    }
    return res
}