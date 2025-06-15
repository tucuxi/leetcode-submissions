func findSafeWalk(grid [][]int, health int) bool {
    m, n := len(grid), len(grid[0])
    v := make([][]int, m)
    for i := range v {
        v[i] = make([]int, n)
    }
    steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for q := [][]int{{0, 0, health}}; len(q) > 0; q = q[1:] {
        r, c := q[0][0], q[0][1]
        h := q[0][2] - grid[r][c]
        if h <= v[r][c] {
            continue
        }
        v[r][c] = h
        if r == m-1 && c == n-1 && h > 0 {
            return true
        }
        for _, s := range steps {
            r2, c2 := r + s[0], c + s[1]
            if r2 >= 0 && r2 < m && c2 >= 0 && c2 < n {
                q = append(q, []int{r2, c2, h})            
            }
        }
    }   
    return false
}