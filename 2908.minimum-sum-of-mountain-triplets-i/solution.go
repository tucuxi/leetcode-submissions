func minimumSum(nums []int) int {
    rmin := func() []int {
        res := make([]int, len(nums))
        p := math.MaxInt
        for i := len(nums) - 1; i >= 0; i-- {
            res[i] = p 
            if nums[i] < p {
                p = nums[i]
            }
        }
        return res
    }()
    return func() int {
        lmin := nums[0]
        res := math.MaxInt
        for i := 1; i < len(nums) - 1; i++ {
            if lmin < nums[i] && rmin[i] < nums[i] {
                if s := lmin + nums[i] + rmin[i]; s < res {
                    res = s
                }
            }
            if nums[i] < lmin {
                lmin = nums[i]
            }
        }
        if res == math.MaxInt {
            return -1
        }
        return res
    }()
}