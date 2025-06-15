func jobScheduling(startTime []int, endTime []int, profit []int) int {
    idx := make([]int, len(startTime))
    for i := range idx {
        idx[i] = i
    }
    sort.Slice(idx, func(i, j int) bool {
        return endTime[idx[i]] < endTime[idx[j]]
    })
    dp := make([]int, len(startTime))
    dp[0] = profit[idx[0]]
    for i := 1; i < len(dp); i++ {
        j := sort.Search(i, func(k int) bool {
            return endTime[idx[k]] > startTime[idx[i]]
        })
        if j > 0 {
            dp[i] = max(dp[i-1], profit[idx[i]] + dp[j-1])
        } else {
            dp[i] = max(dp[i-1], profit[idx[i]])
        }
    }
    return dp[len(dp) - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


