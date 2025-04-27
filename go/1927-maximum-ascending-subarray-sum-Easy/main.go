func maxAscendingSum(nums []int) int {
    sum := nums[0]
    maxSum := sum
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i - 1] {
            sum = nums[i]
        } else {
            sum += nums[i]
        }
        maxSum = max(maxSum, sum)
    }
    return maxSum
}