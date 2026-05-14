func findValidElements(nums []int) []int {
    r := func() []int {
        r := make([]int, len(nums))
        i := len(nums)-1
        r[i] = math.MinInt
        for ; i > 0; i-- {
            r[i-1] = max(nums[i], r[i])
        }
        return r
    }()

    l := math.MinInt
    
    res := make([]int, 0)

    for i, n := range nums {
        if n > l || n > r[i] {
            res = append(res, n)
        }
        l = max(l, n)
    }

    return res
}