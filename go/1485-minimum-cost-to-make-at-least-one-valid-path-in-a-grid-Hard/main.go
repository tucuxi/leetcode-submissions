func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    best := make([][]int, m)
    for i := range m {
        best[i] = make([]int, n)
        for j := range n {
            best[i][j] = math.MaxInt
        }
    }
    best[0][0] = 0

    move := [][]int{
        {},
        { 0,  1},
        { 0, -1},
        { 1,  0},
        {-1,  0},
    }

    for q := [][3]int{{0, 0, 0}}; len(q) > 0; q = q[1:] {
        i, j, c := q[0][0], q[0][1], q[0][2]
        g := grid[i][j]

        for k := 1; k < len(move); k++ {
            i2 := i + move[k][0]
            j2 := j + move[k][1]
            if i2 >= 0 && i2 < m && j2 >= 0 && j2 < n {
                c2 := c
                if g != k {
                    c2++
                }
                if c2 < best[i2][j2] {
                    best[i2][j2] = c2
                    q = append(q, [3]int{i2, j2, c2})
                }
            }
        }
    }

    return best[m - 1][n - 1]
}