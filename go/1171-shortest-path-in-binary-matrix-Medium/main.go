const (
    empty = 0
    visited = 2
)

func shortestPathBinaryMatrix(grid [][]int) int {
    last := len(grid) - 1
    length := 0
    queue := [][2]int{{0, 0}}
    for len(queue) > 0 {
        length++
        for i := len(queue); i > 0; i-- {
            r, c := queue[0][0], queue[0][1]
            queue = queue[1:]
            if grid[r][c] != empty {
                continue
            }
            if r == last && c == last {
                return length
            }
            grid[r][c] = visited
            for rr := max(r - 1, 0); rr <= min(r + 1, last); rr++ {
                for cc := max(c - 1, 0); cc <= min(c + 1, last); cc++ {
                    if grid[rr][cc] == empty {
                        queue = append(queue, [2]int{rr, cc})
                    }
                }
            }
        }
    }
    return -1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}