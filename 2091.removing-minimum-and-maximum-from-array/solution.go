func minimumDeletions(nums []int) int {
    minIndex, maxIndex := 0, 0
    for i := range nums {
        if nums[i] < nums[minIndex] {
            minIndex = i
        }
        if nums[i] > nums[maxIndex] {
            maxIndex = i
        }
    }

    return min(
        max(minIndex, maxIndex) + 1,
        minIndex + 1 + len(nums) - maxIndex,
        len(nums) - minIndex + maxIndex + 1,
        len(nums) - min(minIndex, maxIndex),
    )
}