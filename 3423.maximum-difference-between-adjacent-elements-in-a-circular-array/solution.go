func maxAdjacentDistance(nums []int) int {
    previous := nums[len(nums) - 1]
    maxDiff := 0
    for _, num := range nums {
        maxDiff = max(maxDiff, num - previous, previous - num)
        previous = num
    }
    return maxDiff
}