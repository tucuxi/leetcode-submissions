func smallestRangeII(nums []int, k int) int {
    sort.Ints(nums)
    res := nums[len(nums) - 1] - nums[0]
    for i := 1; i < len(nums); i++ {
        hi := max(nums[len(nums) - 1] - k, nums[i-1] + k)
        lo := min(nums[0] + k, nums[i] - k)
        res = min(res, hi - lo)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}