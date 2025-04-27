const modulo = 1_000_000_007

func countPaths(grid [][]int) int {
    memo := make([][]int, len(grid))
    for r := range memo {
        memo[r] = make([]int, len(grid[r]))
    }
    
    dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
    
    var dfs func(int, int) int
    dfs = func(r, c int) int {
        if memo[r][c] > 0 {
            return memo[r][c]
        }   
        res := 1
        for _, d := range dir {
            r2, c2 := r + d[0], c + d[1]
            if 0 <= r2 && r2 < len(grid) && 0 <= c2 && c2 < len(grid[r2]) && grid[r2][c2] < grid[r][c] {
                res = (res + dfs(r2, c2)) % modulo
            }
        }
        memo[r][c] = res
        return res
    }
    
    count := 0
    for r := range grid {
        for c := range grid[r] {
            count = (count + dfs(r, c)) % modulo
        }
    }    
    return count
}