func getMaximumGold(grid [][]int) int {
    res := 0
    cur := 0

    var dfs func(int, int)
    dfs = func(r, c int) {
        if g := grid[r][c]; g > 0 {
            grid[r][c] = 0
            cur += g
            res = max(res, cur)
            if r > 0 {
                dfs(r-1, c)
            }
            if r < len(grid) - 1 {
                dfs(r+1, c)
            }
            if c > 0 {
                dfs(r, c-1)
            }
            if c < len(grid[0]) - 1 {
                dfs(r, c+1)
            }
            grid[r][c] = g
            cur -= g
        }
    }

    for r := range grid {
        for c := range grid[0] {
            dfs(r, c)
        }
    }

    return res
}