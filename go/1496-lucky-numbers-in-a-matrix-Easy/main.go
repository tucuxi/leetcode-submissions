func luckyNumbers (matrix [][]int) []int {
    minRow := []int{}
    for y := range matrix {
        minIdx := 0
        for x := range matrix[y] {
            if matrix[y][x] < matrix[y][minIdx] {
                minIdx = x
            }
        }
        minRow = append(minRow, minIdx)
    }
    res := []int{}
    for y, x := range minRow {
        maxVal := 0
        for yy := range matrix {
            if matrix[yy][x] > maxVal {
                maxVal = matrix[yy][x]
            }
        }
        if maxVal == matrix[y][x] {
            res = append(res, maxVal)
        }
    }
    return res
}