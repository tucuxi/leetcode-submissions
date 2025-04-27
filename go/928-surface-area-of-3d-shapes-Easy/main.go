func surfaceArea(grid [][]int) int {
    a := 0
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] > 0 {
                a += 4 * grid[i][j] + 2
                if i > 0 {
                    a -= min(grid[i-1][j], grid[i][j]) * 2
                }
                if j > 0 {
                    a -= min(grid[i][j-1], grid[i][j]) * 2
                }
            }
        }
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}