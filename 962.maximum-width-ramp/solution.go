func maxWidthRamp(nums []int) int {
    indices := make([]int, len(nums))
    for i := range indices {
        indices[i] = i
    }
    slices.SortStableFunc(indices, func(a, b int) int { return nums[a] - nums[b] })
    minIndex := len(nums)
    maxWidth := 0
    for _, index := range indices {
        maxWidth = max(index - minIndex, maxWidth)
        minIndex = min(index, minIndex)
    }
    return maxWidth
}