func arithmeticTriplets(nums []int, diff int) int {
    n := len(nums)
    res := 0
    for i := 0; i < n; i++ {
        loop:
        for j := i+1; j < n; j++ {
            if nums[j] - nums[i] == diff {
                for k := j+1; k < n; k++ {
                    if nums[k] - nums[j] == diff {
                        res++
                        break loop
                    }
                }
            }
        }
    }
    return res
}