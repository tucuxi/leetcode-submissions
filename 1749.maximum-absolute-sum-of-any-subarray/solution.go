func maxAbsoluteSum(nums []int) int {
    minSum, maxSum, sum := 0, 0, 0
    for _, num := range nums {
        sum += num
        minSum = min(minSum, sum)
        maxSum = max(maxSum, sum)
    }
    return maxSum - minSum
}