func maxAreaOfIsland(grid [][]int) int {
    max := 0
    
    var dfs func(int, int, int) int
    dfs = func(r int, c int, a int) int {
        if grid[r][c] != 1 {
            return a
        }
        grid[r][c] = 2
        a++
        if r > 0 {
            a = dfs(r-1, c, a)
        }
        if c > 0 {
            a = dfs(r, c-1, a)
        }
        if r+1 < len(grid) {
            a = dfs(r+1, c, a)
        } 
        if c+1 < len(grid[r]) {
            a = dfs(r, c+1, a)
        }
        return a
    }
    
    for i := range grid {
        for j := range grid[i] {
            if a := dfs(i, j, 0); a > max {
                max = a
            }
        }
    }
    return max
}