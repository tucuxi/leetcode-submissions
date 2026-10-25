func peakValleyDifference(nums []int) int {
    res := 0
    for i := 1; i < len(nums) - 1; i++ {
        if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
            res += nums[i]
        } else if nums[i-1] > nums[i] && nums[i] < nums[i+1] {
            res -= nums[i]
        }
    } 
    return res
}