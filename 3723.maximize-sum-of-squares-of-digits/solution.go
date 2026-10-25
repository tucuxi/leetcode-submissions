func maxSumOfSquares(num int, sum int) string {
    result := make([]byte, num)
    rest := sum

    for i := range num {
        digit := min(rest, 9)
        rest -= digit
        result[i] = byte('0' + digit)
    }

    if rest > 0 {
        return ""
    }
    return string(result)
}