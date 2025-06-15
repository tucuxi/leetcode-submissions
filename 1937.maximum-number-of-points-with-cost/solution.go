func maxPoints(points [][]int) int64 {
    m := len(points)
    n := len(points[0])
    previousRow := make([]int64, n)
    currentRow := make([]int64, n)

    for j := range n {
        previousRow[j] = int64(points[0][j])
    }
    for i := 1; i < m; i++ {
        var leftMax int64
        for j := range n {
            leftMax = max(leftMax - 1, previousRow[j])
            currentRow[j] = leftMax
        }
        var rightMax int64
        for j := n-1; j >= 0; j-- {
            rightMax = max(rightMax - 1, previousRow[j])
            currentRow[j] = max(currentRow[j], rightMax) + int64(points[i][j])
        }
        previousRow, currentRow = currentRow, previousRow
    }

    var res int64
    for j := range n {
        res = max(res, previousRow[j])
    }
    return res
}