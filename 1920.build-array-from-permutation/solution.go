func buildArray(nums []int) []int {
    ans := make([]int, len(nums))
    for i := range ans {
        ans[i] = nums[nums[i]]
    }
    return ans
}