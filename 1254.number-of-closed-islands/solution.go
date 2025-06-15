func closedIsland(grid [][]int) int {
    var dfs func(int, int) bool
    dfs = func(r, c int) bool {
        if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
            return false
        }
        if grid[r][c] > 0 {
            return true
        }
        grid[r][c] = 2
        up := dfs(r - 1, c)
        down := dfs(r + 1, c)
        left := dfs(r, c - 1)
        right := dfs(r, c + 1)
        return up && down && left && right
    }
    count := 0
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 0 && dfs(r, c) {
                count++
            }
        }
    }
    return count
}