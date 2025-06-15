func duplicateNumbersXOR(nums []int) int {
    var h [51]bool
    res := 0

    for _, n := range nums {
        if h[n] {
            res ^= n
        } else {
            h[n] = true
        }
    }
    return res
}