func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })
    max := 0
    for i := 1; i < len(points); i++ {
        if dx := points[i][0] - points[i - 1][0]; dx > max {
            max = dx
        } 
    }
    return max
}