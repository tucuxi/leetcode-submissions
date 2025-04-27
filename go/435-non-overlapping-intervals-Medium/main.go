func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    end := intervals[0][1]
    res := 0
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < end {
            res++
        } else {
            end = intervals[i][1]
        }
    }
    return res
}