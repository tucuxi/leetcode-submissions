func findMaxFish(grid [][]int) int {
    fish := 0

    var dfs func(int, int) int

    dfs = func(r, c int) int {
        if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] == 0 {
            return 0
        }
        f := grid[r][c]
        grid[r][c] = 0
        for _, next := range [][]int{{r-1, c}, {r+1, c}, {r, c-1}, {r, c+1}} {
            f += dfs(next[0], next[1])
        }
        return f
    }

    for r := range grid {
        for c := range grid[r] {
            fish = max(fish, dfs(r, c))
        }
    }
    return fish
}