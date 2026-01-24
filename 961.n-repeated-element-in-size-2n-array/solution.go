func repeatedNTimes(nums []int) int {
    for i := 2; i < len(nums); i++ {
        if (nums[i - 2] == nums[i] || nums[i - 1] == nums[i]) {
            return nums[i]
        }
    }
    return nums[0]
}