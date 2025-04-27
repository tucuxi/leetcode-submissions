func shortestBridge(grid [][]int) int {
    visited := make([][]bool, len(grid))
    for row := range grid {
        visited[row] = make([]bool, len(grid[row]))
    }
    dir := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
    queue := make([][]int, 0)
    
    var dfs func(int, int)
    dfs = func(row, col int) {
        if grid[row][col] == 1 && !visited[row][col] {
            visited[row][col] = true
            queue = append(queue, []int{row, col, 0})
            for _, d := range dir {
                r, c := row + d[0], col + d[1]
                if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) {
                    dfs(r, c)
                }
            }
        }
    }

    outer:
    for row := range grid {
        for col := range grid[row] {
            if grid[row][col] == 1 {
                dfs(row, col)
                break outer
            }
        }
    }
    
    for len(queue) > 0 {
        row, col, dist := queue[0][0], queue[0][1], queue[0][2]
        queue = queue[1:]        
        for _, d := range dir {
            r, c := row + d[0], col + d[1]
            if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) && !visited[r][c] {
                if grid[r][c] == 1 {
                    return dist
                }
                visited[r][c] = true
                queue = append(queue, []int{r, c, dist + 1})
            }
        }
    }
        
    panic("no solution")
}