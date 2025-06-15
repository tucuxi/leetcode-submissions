func maxSubArray(nums []int) int {
    running, largest := 0, math.MinInt
    for _, n := range nums {
        running += n
        largest = max(largest, running)
        running = max(running, 0)
    }
    return largest
}