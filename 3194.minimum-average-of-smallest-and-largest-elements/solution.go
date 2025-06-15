func minimumAverage(nums []int) float64 {
    slices.Sort(nums)
    minAverage := math.MaxFloat64
    for i := 0; i < len(nums) / 2; i++ {
        average := (float64(nums[i]) + float64(nums[len(nums) - 1 - i])) / 2
        if average < minAverage {
            minAverage = average
        }
    }
    return minAverage
}