func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)

    gap := func(i int) int {
        switch i {
            case 0: return startTime[0]
            case n: return eventTime - endTime[n - 1]
            default: return startTime[i] - endTime[i - 1]
        }
    }

    largestRightGap := func() []int {
        r := make([]int, n + 1)
        r[n - 1] = gap(n)
        for i := n - 1; i > 0; i-- {
            r[i - 1] = max(r[i], gap(i))
        }
        return r
    }()

    largestLeftGap := 0
    maxFreeTime := 0

    for i := range n {
        meetingDuration := endTime[i] - startTime[i]
        freeTime := gap(i) + gap(i + 1)
        if meetingDuration <= largestLeftGap || meetingDuration <= largestRightGap[i + 1] {
            freeTime += meetingDuration
        }
        maxFreeTime = max(maxFreeTime, freeTime)
        largestLeftGap = max(largestLeftGap, gap(i))
    }

    return maxFreeTime
}