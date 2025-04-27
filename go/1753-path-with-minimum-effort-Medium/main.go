const maxHeight = 1_000_000

func minimumEffortPath(heights [][]int) int {
    return sort.Search(maxHeight, func(k int) bool {
        return solve(heights, k)
    })    
}

func solve(heights [][]int, k int) bool {
    visited := make([][]bool, len(heights))
    for i := range visited {
        visited[i] = make([]bool, len(heights[i]))
    }
    
    var dfs func(int, int) bool
    dfs = func(r, c int) bool {
        visited[r][c] = true        
        re := len(heights) - 1
        ce := len(heights[0]) - 1
        if r == re && c == ce {
            return true
        }
        for _, s := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            r2 := r + s[0]
            c2 := c + s[1]
            if r2 >= 0 && r2 <= re && c2 >= 0 && c2 <= ce && !visited[r2][c2] && abs(heights[r][c] - heights[r2][c2]) <= k && dfs(r2, c2) {
                return true
            }
        }
        return false
    }
    
    return dfs(0, 0)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}