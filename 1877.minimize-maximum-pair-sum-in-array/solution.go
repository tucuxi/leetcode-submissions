func minPairSum(nums []int) int {
    slices.Sort(nums)
    last := len(nums) - 1
    res := 0
    for i := range len(nums) / 2 {
        res = max(res, nums[i] + nums[last - i])
    }
    return res
}