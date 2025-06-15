func longestSubarray(nums []int) int {
    maxBitwise := 0
    maxLength := 0
    curLength := 0
    for _, n := range nums {
        switch {
        case n > maxBitwise:
            maxBitwise = n
            maxLength = 1
            curLength = 1
        case n == maxBitwise:
            curLength++
            if curLength > maxLength {
                maxLength = curLength
            }
        default:
            curLength = 0
        }
    }
    return maxLength
}