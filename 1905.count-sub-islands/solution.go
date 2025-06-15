func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    var (
        res = 0
        island = 1
        dfs func(int, int) bool
    )
    
    dfs = func(i, j int) bool {
        grid2[i][j] = island
        res := grid1[i][j] == 1
        for _, step := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            i2, j2 := i + step[0], j + step[1]
            if i2 >= 0 && i2 < len(grid2) && j2 >= 0 && j2 < len(grid2[i2]) && grid2[i2][j2] == 1 {
                if !dfs(i2, j2) {
                    res = false
                }
            }
        }
        return res
    }
    
    
    for i := range grid2 {
        for j := range grid2[i] {
            if grid2[i][j] == 1 {
                island++
                if dfs(i, j) {
                    res++
                }
            }
        }
    }
    
    return res
}

