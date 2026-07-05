func maxDigitRange(nums []int) int {
    sum := 0
    maxRange := 0
    for _, num := range nums {
        dr := digitRange(num)
        if dr > maxRange {
            sum = num
            maxRange = dr
        } else if dr == maxRange {
            sum += num
        }
    }
    return sum
}

func digitRange(n int) int {
    minDigit, maxDigit := 9, 0
    for i := n; i > 0; i /= 10 {
        digit := i % 10
        minDigit = min(digit, minDigit)
        maxDigit = max(digit, maxDigit)
    }
    return maxDigit - minDigit
}