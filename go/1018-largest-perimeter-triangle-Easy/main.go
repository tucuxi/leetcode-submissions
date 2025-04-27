func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    for i := len(nums) - 3; i >= 0; i-- {
        if nums[i + 2] < nums[i + 1] + nums[i] {
            return nums[i] + nums[i + 1] + nums[i + 2]
        }
    }
    return 0
}