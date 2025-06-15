func triangularSum(nums []int) int {
    for n := len(nums); n > 1; n-- {
        for i := 0; i < n - 1; i++ {
            nums[i] = (nums[i] + nums[i + 1]) % 10
        }
    }
    return nums[0]
}