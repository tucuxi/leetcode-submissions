func hardestWorker(n int, logs [][]int) int {
    id := math.MaxInt
    maxDuration := math.MinInt
    time := 0
    for i := range logs {
        duration := logs[i][1] - time
        time = logs[i][1]
        if duration == maxDuration {
            if logs[i][0] < id {
                id = logs[i][0]
            }
        } else if duration > maxDuration {
            maxDuration = duration
            id = logs[i][0]
        }
    }
    return id
}