func satisfiesConditions(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    for i := range m {
        for j := range n {
            if i+1 < m && grid[i][j] != grid[i+1][j] {
                return false
            }
            if j+1 < n && grid[i][j] == grid[i][j+1] {
                return false
            }
        }
    }
    return true
}