func countDigitOccurrences(nums []int, digit int) int {
    res := 0
    for _, num := range nums {
        for n := num; n > 0; n /= 10 {
            if n%10 == digit {
                res++
            }
        } 
    }
    return res
}