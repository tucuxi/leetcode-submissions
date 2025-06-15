func maxSubarraySumCircular(nums []int) int {
    curMin := nums[0]
    curMax := nums[0]
    minSum := nums[0]
    maxSum := nums[0]
    total := nums[0]
    for _, n := range nums[1:] {
        curMax = max(curMax + n, n)
        maxSum = max(maxSum, curMax)
        curMin = min(curMin + n, n)
        minSum = min(minSum, curMin)
        total += n
    }
    if maxSum > 0 {
        return max(maxSum, total - minSum)
    }
    return maxSum
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}