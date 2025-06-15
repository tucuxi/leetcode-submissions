func maxSubsequence(nums []int, k int) []int {
    idx := make([]int, len(nums))
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i1, i2 int) bool {
        return nums[idx[i1]] < nums[idx[i2]]
    })
    top := idx[len(nums) - k:]
    sort.Ints(top)
    res := make([]int, len(top))
    for i := range top {
        res[i] = nums[top[i]]
    }
    return res
}
