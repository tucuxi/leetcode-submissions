func separateDigits(nums []int) []int {
    res := []int{}
    for _, n := range nums {
        digits := []int{}
        for n > 0 {
            digits = append(digits, n % 10)
            n /= 10
        }
        for i := len(digits) - 1; i >= 0; i-- {
            res = append(res, digits[i])
        }
    }
    return res
}