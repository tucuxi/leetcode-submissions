func countSubarrays(nums []int) int {
    count := 0
    for i := 2; i < len(nums); i++ {
        if 2 * (nums[i - 2] + nums[i]) == nums[i - 1] {
            count++
        }
    }
    return count
}