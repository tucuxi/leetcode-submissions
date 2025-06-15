func checkValidGrid(grid [][]int) bool {
    steps := [...][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
    n := len(grid)
    r, c := 0, 0
    if grid[r][c] != 0 {
        return false
    }
    
    outer:
    for i := 1; i < n * n; i++ {
        for _, s := range steps {
            r2, c2 := r + s[0], c + s[1]
            if r2 >= 0 && r2 < n && c2 >= 0 && c2 < n && grid[r2][c2] == i {
                r, c = r2, c2
                continue outer
            }
        }
        return false
    }
    return true
}