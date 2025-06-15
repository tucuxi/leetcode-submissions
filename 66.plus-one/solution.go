func plusOne(digits []int) []int {
    carry := 1
    for i := len(digits) - 1; i >= 0; i-- {
        digit := digits[i] + carry
        digits[i], carry = digit % 10, digit / 10
    }
    if carry > 0 {
        return append([]int{carry}, digits...)
    }
    return digits
}