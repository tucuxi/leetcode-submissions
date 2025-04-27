func minimumPairRemoval(nums []int) int {
    for operations := 0; ; operations++ {
        nonDecreasing := true
        minSum := math.MaxInt
        minSumIndex := -1
        for i := 1; i < len(nums); i++ {
            if nums[i] < nums[i - 1] {
                nonDecreasing = false
            }
            if sum := nums[i - 1] + nums[i]; sum < minSum {
                minSum = sum
                minSumIndex = i
            }
        }
        if nonDecreasing {
            return operations
        }
        prefix := nums[: minSumIndex - 1]
        suffix := nums[minSumIndex + 1 :]
        nums = append(prefix, minSum)
        nums = append(nums, suffix...)
    }
}