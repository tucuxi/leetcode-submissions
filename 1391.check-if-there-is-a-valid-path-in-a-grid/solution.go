func hasValidPath(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    
    v := make([][]bool, m)
    for i := range v {
        v[i] = make([]bool, n)
    }

    neighbors := func(r, c int) [][]int {
        var res [][]int
        g := grid[r][c]

        if r > 0 {
            if g == 2 || g == 5 || g == 6 {
                h := grid[r-1][c]
                if h == 2 || h == 3 || h == 4 {
                    res = append(res, []int{r-1, c})
                }
            }    
        }
        if c > 0 {
            if g == 1 || g == 3 || g == 5 {
                h := grid[r][c-1]
                if h == 1 || h == 4 || h == 6 {
                    res = append(res, []int{r, c-1})
                }
            }
        }
        if r < m-1 {
            if g == 2 || g == 3 || g == 4 {
                h := grid[r+1][c]
                if h == 2 || h == 5 || h == 6 {
                    res = append(res, []int{r+1, c})
                }
            }
        }
        if c < n-1 {
            if g == 1 || g == 4 || g == 6 {
                h := grid[r][c+1]
                if h == 1 || h == 3 || h == 5 {
                    res = append(res, []int{r, c+1})
                }
            }
        }
        return res
    }

    var dfs func(r, c int) bool

    dfs = func(r, c int) bool {
        if r == m-1 && c == n-1 {
            return true
        }
        v[r][c] = true
        for _, nb := range neighbors(r, c) {
            if !v[nb[0]][nb[1]] {
                if dfs(nb[0], nb[1]) {
                    return true
                }
            }
        }
        return false
    }

    return dfs(0, 0)
}