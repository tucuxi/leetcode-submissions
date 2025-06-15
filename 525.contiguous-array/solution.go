func findMaxLength(nums []int) int {
    h := make(map[int]int)
    h[0] = -1
    sum := 0
    maxLength := 0
    for i, n := range nums {
        if n > 0 {
            sum++
        } else {
            sum--
        }
        j, exists := h[sum]
        if exists {
            maxLength = max(maxLength, i-j)
        } else {
            h[sum] = i
        }
    }
    return maxLength
}