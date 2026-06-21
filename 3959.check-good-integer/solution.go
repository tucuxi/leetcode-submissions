func checkGoodInteger(n int) bool {
    digitSum := 0
    squareSum := 0
    for r := n; r > 0; r /= 10 {
        digit := r % 10
        digitSum += digit
        squareSum += digit * digit
    }
    return squareSum - digitSum >= 50
}