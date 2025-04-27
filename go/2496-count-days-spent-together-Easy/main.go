func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
    a1 := dayOfYear(arriveAlice)
    a2 := dayOfYear(leaveAlice)
    b1 := dayOfYear(arriveBob)
    b2 := dayOfYear(leaveBob)
    p := min(a2, b2) - max(a1, b1)
    if p < 0 {
        return 0
    }
    return p + 1
}

func dayOfYear(s string) int {
    var month, day, md int
    fmt.Sscanf(s, "%d-%d", &month, &day)
    switch month {
    case 2: md = 31
    case 3: md = 59
    case 4: md = 90
    case 5: md = 120
    case 6: md = 151
    case 7: md = 181
    case 8: md = 212
    case 9: md = 243
    case 10: md = 273
    case 11: md = 304
    case 12: md = 334
    }
    return md + day
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}