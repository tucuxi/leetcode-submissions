func findValueOfPartition(nums []int) int {
    value := math.MaxInt
    slices.Sort(nums) 
    for i := 1; i < len(nums); i++ {
        value = min(value, nums[i] - nums[i - 1])
    }
    return value
}