func minimumDifference(nums []int, k int) int {
    slices.Sort(nums)
    diff := nums[k - 1] - nums[0]
    for i := k; i < len(nums); i++ {
        diff = min(diff, nums[i] - nums[i - k + 1]) 
    }
    return diff
}