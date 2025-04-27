func nearestExit(maze [][]byte, entrance []int) int {
    m, n := len(maze), len(maze[0])
    q, d := [][]int{entrance}, 0
    maze[entrance[0]][entrance[1]] = '-'
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            row, col := q[0][0], q[0][1]
            q = q[1:]
            for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
                nrow, ncol := row + dir[0], col + dir[1]
                if nrow < 0 || nrow == m || ncol < 0 || ncol == n {
                    if row != entrance[0] || col != entrance[1] {
                        return d
                    }
                } else if maze[nrow][ncol] == '.' {
                    maze[nrow][ncol] = '-'
                    q = append(q, []int{nrow, ncol})
                }
            }
        }
        d++
    }
    return -1
}