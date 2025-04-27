func convertTime(current string, correct string) int {
    var currentHours, currentMinutes, correctHours, correctMinutes int
    fmt.Sscanf(current, "%2d:%2d", &currentHours, &currentMinutes)
    fmt.Sscanf(correct, "%2d:%2d", &correctHours, &correctMinutes)
    d := (correctHours - currentHours) * 60 + correctMinutes - currentMinutes
    res := 0
    for _, k := range []int{60, 15, 5, 1} {
        res += d / k
        d %= k
    }
    return res
}