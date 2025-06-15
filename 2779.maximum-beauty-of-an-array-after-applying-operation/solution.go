func maximumBeauty(nums []int, k int) int {
    slices.Sort(nums)
    res := 0
    q := 0
    for p := range nums {
        for q < len(nums) && nums[q] - nums[p] <= 2 * k {
            q++
        }
        res = max(res, q - p)
    }
    return res
}