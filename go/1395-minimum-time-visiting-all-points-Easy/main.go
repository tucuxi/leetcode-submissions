func minTimeToVisitAllPoints(points [][]int) int {
    res := 0
    for i := 1; i < len(points); i++ {
        dx := abs(points[i][0] - points[i-1][0])
        dy := abs(points[i][1] - points[i-1][1])
        if dx > dy {
            res += dx
        } else {
            res += dy
        }
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}