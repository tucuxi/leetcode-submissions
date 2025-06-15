func sumOfEncryptedInt(nums []int) int {
    res := 0
    for _, n := range nums {
        d, l := largestDigitAndLength(n)
        e := 0
        for ; l > 0; l-- {
            e = 10 * e + d
        }
        res += e
    }    
    return res
}

func largestDigitAndLength(n int) (int, int) {
    length := 0
    largestDigit := 0
    for ; n > 0; n /= 10 {
        digit := n % 10
        largestDigit = max(largestDigit, digit)
        length++
    }
    return largestDigit, length
}