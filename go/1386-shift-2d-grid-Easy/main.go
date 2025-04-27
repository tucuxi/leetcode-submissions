func shiftGrid(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    for ; k > 0; k-- {
        h := grid[m - 1][n - 1]
        for i := m - 1; i >= 0; i-- {
            for j := n - 1; j >= 1; j-- {
                grid[i][j] = grid[i][j - 1]
            }
            if i > 0 {
                grid[i][0] = grid[i - 1][n - 1]
            } else {
                grid[i][0] = h
            }
        }
    }
    return grid
}