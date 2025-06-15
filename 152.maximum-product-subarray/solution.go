func maxProduct(nums []int) int {
    res := nums[0]
    minimum := nums[0]
    maximum := nums[0]
    for _, n := range nums[1:] {
        if n < 0 {
            minimum, maximum = maximum, minimum
        }
        minimum = min(n, n * minimum)
        maximum = max(n, n * maximum)
        res = max(res, maximum)
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}