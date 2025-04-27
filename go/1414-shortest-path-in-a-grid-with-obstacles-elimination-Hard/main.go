type step struct{row, col, used, cnt int}

func shortestPath(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    visited := make([][]int, m)
    for i := range visited {
        visited[i] = make([]int, n)
        for j := range visited[i] {
            visited[i][j] = math.MaxInt
        }
    }
    q := []step{{}}
    for len(q) > 0 {
        c := q[0]
        q = q[1:]
        if c.row == m-1 && c.col == n-1 {
            return c.cnt
        }
        for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0 ,1}} {
            newRow, newCol := c.row + dir[0], c.col + dir[1]
            if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n {
                continue
            }
            newUsed := c.used + grid[newRow][newCol]
            if newUsed > k || newUsed >= visited[newRow][newCol] {
                continue
            }
            visited[newRow][newCol] = newUsed
            q = append(q, step{newRow, newCol, newUsed, c.cnt + 1})
        }
    }
    return -1
}