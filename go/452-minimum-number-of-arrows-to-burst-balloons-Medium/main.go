func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] })
    res := 1
    end := points[0][1]
    for i := 1; i < len(points); i++ {
        if points[i][0] > end {
            end = points[i][1]
            res++
        }
    }
    return res
}