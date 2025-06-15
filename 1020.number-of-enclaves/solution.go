func numEnclaves(grid [][]int) int {
   
    var dfs func(int, int, int) (int, bool)
    dfs = func(r, c, s int) (int, bool) {
        if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
            return s, false
        }
        if grid[r][c] == 0 || grid[r][c] == 2 {
            return s, true
        }
        grid[r][c] = 2
        sl, el := dfs(r, c - 1, 0)
        sr, er := dfs(r, c + 1, 0)
        su, eu := dfs(r - 1, c, 0)
        sd, ed := dfs(r + 1, c, 0)
        return s + sl + sr + su + sd + 1, el && er && eu && ed
    }
    
    enclaves := 0
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 1 {
                if s, e := dfs(r, c, 0); e {
                    enclaves += s
                }
            }
        }
    }
    
    return enclaves
}