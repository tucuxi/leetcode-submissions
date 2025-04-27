func maximumScore(nums []int, multipliers []int) int {
    n := len(nums)
    m := len(multipliers)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }
    for k := m-1; k >= 0; k-- {
        for i := k; i >= 0; i-- {
            j := n - 1 - (k - i)
            l := nums[i] * multipliers[k] + dp[i + 1][k + 1]
            r := nums[j] * multipliers[k] + dp[i][k + 1]
            if l < r {
                dp[i][k] = r
            } else {
                dp[i][k] = l
            }
        }
    }
    return dp[0][0]
}