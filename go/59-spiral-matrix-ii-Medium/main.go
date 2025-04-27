func generateMatrix(n int) [][]int {
    res := make([][]int, n)
    for i := range res {
        res[i] = make([]int, n)
    }
    row, col := 0, 0
    dr, dc := 0, 1
    for k := 1; k <= n * n; k++ {
        res[row][col] = k
        if r2, c2 := row + dr, col + dc; !inRange(r2, 0, n) || !inRange(c2, 0, n) || res[r2][c2] != 0 {
            dr, dc = dc, -dr
        }
        row += dr
        col += dc
    }
    return res
}

func inRange(a, b, c int) bool {
    return a >= b && a < c
}