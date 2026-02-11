func minimumOperations(nums []int) int {
    dp := make([]int, 4)
    
    for _, n := range nums {
        dp[n] = slices.Max(dp[:n+1]) + 1
    }

    return len(nums) - slices.Max(dp)
}