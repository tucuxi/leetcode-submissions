func findMinDifference(timePoints []string) int {
    minutes := make([]int, len(timePoints))
    for i := range timePoints {
        var hh, mm int
        fmt.Sscanf(timePoints[i], "%2d:%2d", &hh, &mm)
        minutes[i] = hh * 60 + mm
    }
    sort.Ints(minutes)
    res := 24 * 60 - minutes[len(minutes) - 1] + minutes[0]
    for i := 1; i < len(minutes); i++ {
        res = min(res, minutes[i] - minutes[i - 1])
    }
    return res
}