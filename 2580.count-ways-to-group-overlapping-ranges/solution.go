func countWays(ranges [][]int) int {
    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i][0] < ranges[j][0] || ranges[i][0] == ranges[j][0] && ranges[i][1] > ranges[j][1]
    })
    res := 1
    currentEnd := -1
    for i := range ranges {
        if ranges[i][0] > currentEnd {
            res = 2 * res % 1_000_000_007
        }
        if ranges[i][1] > currentEnd {
            currentEnd = ranges[i][1]
        }
    }
    return res
}