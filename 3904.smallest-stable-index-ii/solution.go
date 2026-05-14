func firstStableIndex(nums []int, k int) int {
    n := len(nums)
    
    mins := make([]int, n)
    mins[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        mins[i] = min(nums[i], mins[i+1])
    }

    runningMax := 0
    for i, num := range nums {
        runningMax = max(runningMax, num)
        if runningMax - mins[i] <= k {
            return i
        }
    }

    return -1
}