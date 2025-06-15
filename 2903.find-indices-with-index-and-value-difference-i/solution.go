func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    for i := range nums {
        for j := i + indexDifference; j < len(nums); j++ {
            if abs(nums[i] - nums[j]) >= valueDifference {
                return []int{i, j}
            }
        }
    }
    return []int{-1, -1}
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}