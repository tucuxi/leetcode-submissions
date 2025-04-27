func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    for r, c := m-1, 0; r >= 0 && c < n;  {
        if matrix[r][c] == target {
            return true
        } else if matrix[r][c] > target {
            r--
        } else {
            c++
        }
    }
    return false
}