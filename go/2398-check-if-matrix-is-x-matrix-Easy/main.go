func checkXMatrix(grid [][]int) bool {
    n := len(grid)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if (i == j || i + j == n - 1) == (grid[i][j] == 0) {
                return false
            }
        }
    }
    return true
}