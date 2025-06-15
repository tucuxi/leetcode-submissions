func minimizeArrayValue(nums []int) int {
    m := 0
    for _, n := range nums {
        if m < n {
            m = n
        }
    }
    return sort.Search(m, func(n int) bool {
        d := 0
        for i := len(nums) - 1; i >= 0; i-- {
            d = max(0, d + nums[i] - n)
        }
        return d == 0
    })
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}