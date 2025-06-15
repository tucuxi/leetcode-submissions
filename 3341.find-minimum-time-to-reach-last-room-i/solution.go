func minTimeToReach(moveTime [][]int) int {
    m, n := len(moveTime), len(moveTime[0])

    dist := make([][]int, m)
    for i := range m {
        dist[i] = make([]int, n)
        for j := range n {
            dist[i][j] = math.MaxInt
        }
    }
    dist[0][0] = 0

    q := [][2]int{{0, 0}}

    improve := func(r1, c1, r2, c2 int) {
        newDist := max(dist[r1][c1], moveTime[r2][c2]) + 1
        if newDist < dist[r2][c2] {
            dist[r2][c2] = newDist
            q = append(q, [2]int{r2, c2})
        }
    }

    for len(q) > 0 {
        r, c := q[0][0], q[0][1]
        q = q[1:]
        if r > 0 {
            improve(r, c, r - 1, c)
        }
        if r + 1 < m {
            improve(r, c, r + 1, c)
        }
        if c > 0 {
            improve(r, c, r, c - 1)
        }
        if c + 1 < n {
            improve(r, c, r, c + 1)
        }
    }

    return dist[m - 1][n - 1]
}