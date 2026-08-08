func maximumProduct(nums []int) int {
    slices.Sort(nums)
    n := len(nums)
    a := nums[n-1] * nums[n-2] * nums[n-3]
    b := nums[n-1] * nums[0] * nums[1]
    return max(a, b)
}