func rotateGrid(grid [][]int, k int) [][]int {
    m := len(grid)
    n := len(grid[0])
    nlayer := min(m / 2, n / 2) 
    
    for layer := 0; layer < nlayer; layer++ {
        r := make([]int, 0)
        c := make([]int, 0)
        val := make([]int, 0)
        for i := layer; i < m - layer - 1; i++ {   // left
            r = append(r, i)
            c = append(c, layer)
            val = append(val, grid[i][layer])
        }
        for j := layer; j < n - layer - 1; j++ {   // down
            r = append(r, m - layer - 1)
            c = append(c, j)
            val = append(val, grid[m-layer - 1][j])
        }
        for i := m - layer - 1; i > layer; i-- {   // right
            r = append(r, i)
            c = append(c, n - layer - 1)
            val = append(val, grid[i][n - layer - 1])
        }
        for j := n - layer - 1; j > layer; j-- {   // up
            r = append(r, layer)
            c = append(c, j)
            val = append(val, grid[layer][j])
        }
        total := len(val)
        kk := k % total
        for i := 0; i < total; i++ {
            idx := (i + total - kk) % total
            grid[r[i]][c[i]] = val[idx]
        }
    }
    return grid
}