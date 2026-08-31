func elevatorRequests(n int, requests []int) int {
    floor := 0
    res := 0

    for _, r := range requests {
        res += abs(r - floor)
        floor = r
    }
    return res
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}