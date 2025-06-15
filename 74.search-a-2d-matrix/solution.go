func searchMatrix(matrix [][]int, target int) bool {
    row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
    if row < 0 {
        return false
    }
    col := sort.Search(len(matrix[row]), func(i int) bool { return matrix[row][i] >= target })
    return col < len(matrix[row]) && matrix[row][col] == target
}