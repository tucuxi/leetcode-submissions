func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    m := max(abs(sx - fx), abs(sy - fy))
    if m == 0 {
        return t != 1
    }
    return t >= m 
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}