func pacificAtlantic(heights [][]int) [][]int {
    m, n := len(heights), len(heights[0])
    atlantic, pacific := make([][]bool, m), make([][]bool, m)
    for r := 0; r < m; r++ {
        atlantic[r] = make([]bool, n)
        atlantic[r][n - 1] = true
        pacific[r] = make([]bool, n)
        pacific[r][0] = true
    }
    for c := 0; c < n; c++ {
        atlantic[m - 1][c] = true
        pacific[0][c] = true
    }
    bfs(heights, atlantic)
    bfs(heights, pacific)
    res := make([][]int, 0, len(heights))
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if atlantic[r][c] && pacific[r][c] {
                res = append(res, []int{r, c})
            }
        }
    }
    return res
}

func bfs(heights [][]int, visited [][]bool) {
    q := make([][]int, 0, len(heights) + len(heights[0]))
    for r := range visited {
        for c := range visited[r] {
            if visited[r][c] {
                q = append(q, []int{r, c})
            }
        }
    }
    for len(q) > 0 {
        r, c := q[0][0], q[0][1]
        q = q[1:]
        for _, nb := range [][]int{{r-1, c}, {r+1,c}, {r, c-1}, {r, c+1}} {
            r1, c1 := nb[0], nb[1]
            if r1 >= 0 && r1 < len(heights) && c1 >= 0 && c1 < len(heights[0]) {
                if !visited[r1][c1] && heights[r1][c1] >= heights[r][c] {
                    q = append(q, nb)
                    visited[r1][c1] = true
                }
            }
        } 
    }
}