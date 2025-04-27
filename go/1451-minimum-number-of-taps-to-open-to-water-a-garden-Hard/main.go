func minTaps(n int, ranges []int) int {
    intervals := make([][2]int, len(ranges))
    for i := range ranges {
        intervals[i][0] = i - ranges[i]
        intervals[i][1] = i + ranges[i]
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    count := 0
    p := 0
    for i := 0; i < len(intervals) - 1; {
        j := i
        for q := p; q < len(intervals) && intervals[q][0] <= i; q++ {
            if intervals[q][1] > j {
                j = intervals[q][1]
                p = q
            }
        }
        if j == i {
            return -1
        }
        count++
        i = j
    }
    return count
}