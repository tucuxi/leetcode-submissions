func largestSumOfAverages(nums []int, k int) float64 {
    n := len(nums)

    p := make([]float64, n+1)
    for i := range n {
        p[i+1] = p[i] + float64(nums[i])
    }

    dp := make([]float64, n)
    for i := range n {
        dp[i] = (p[n]-p[i]) / float64(n-i)
    }

    for range k-1 {
        for i := range n {
            for j := i+1; j < n; j++ {
                dp[i] = max(dp[i], dp[j] + (p[j]-p[i]) / float64(j-i))
            }
        }
    }

    return dp[0]
}