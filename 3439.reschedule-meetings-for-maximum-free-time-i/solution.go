func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    n := len(startTime)

    prefixSum := make([]int, n + 2)
    prefixSum[1] = startTime[0]
    for i := 1; i < n; i++ {
        prefixSum[i + 1] = prefixSum[i] + startTime[i] - endTime[i - 1] 
    }
    prefixSum[n + 1] = prefixSum[n] + eventTime - endTime[n - 1]

    res := 0
    for i := k + 1; i <= n + 1; i++ {
        res = max(res, prefixSum[i] - prefixSum[i - k - 1])
    }
    return res
}

