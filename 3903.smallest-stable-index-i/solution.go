func firstStableIndex(nums []int, k int) int {
    mins := func() []int {
        mins := make([]int, len(nums))
        mn := nums[len(nums) - 1]
        for i := len(nums) - 1; i >= 0; i-- {
            mn = min(mn, nums[i])
            mins[i] = mn
        }
        return mins
    }()

    mx := 0
    for i := range nums {
        mx = max(mx, nums[i])
        if mx - mins[i] <= k {
            return i
        }
    }

    return -1    
}