func maxDistance(grid [][]int) int {
    q := [][2]int{}
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 1 {
                q = append(q, [2]int{r, c} )
            }
        }
    }
    d := -1
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            cell := q[0]
            q = q[1:]
            for _, s := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
                r, c := cell[0] + s[0], cell[1] + s[1]
                if r >= 0 && r < len(grid) && c >= 0 && c < len(grid) && grid[r][c] == 0 {
                    grid[r][c] = 1
                    q = append(q, [2]int{r, c})
                }
             }
        }
        d++
    }
    if d == 0 {
        return -1
    }
    return d
}