func maxAlternatingSum(nums []int) int64 {
    for i := range nums {
        nums[i] *= nums[i]
    }

    slices.Sort(nums)
    
    var res int64

    for i := range len(nums) / 2 {
        res -= int64(nums[i])
    }
    for i := len(nums) / 2; i < len(nums); i++ {
        res += int64(nums[i])
    }
    return res
}