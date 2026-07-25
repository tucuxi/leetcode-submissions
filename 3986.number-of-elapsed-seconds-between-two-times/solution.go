const (
    secondsPerMinute = 60
    secondsPerHour   = 3600
)

func secondsBetweenTimes(startTime string, endTime string) int {
    var sh, sm, ss, eh, em, es int

    fmt.Sscanf(startTime, "%d:%d:%d", &sh, &sm, &ss)
    fmt.Sscanf(endTime, "%d:%d:%d", &eh, &em, &es)

    startSeconds := sh * secondsPerHour + sm * secondsPerMinute + ss
    endSeconds := eh * secondsPerHour + em * secondsPerMinute + es

    return endSeconds - startSeconds
}