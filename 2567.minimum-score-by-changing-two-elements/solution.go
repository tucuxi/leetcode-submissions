func minimizeSum(nums []int) int {
    sort.Ints(nums)
    return min(nums[len(nums) - 3] - nums[0], nums[len(nums) - 2] - nums[1], nums[len(nums) - 1] - nums[2])
}

func min(a, b, c int) int {
    if a < b {
        if a < c {
            return a
        }
        return c
    }
    if b < c {
        return b
    }
    return c
}