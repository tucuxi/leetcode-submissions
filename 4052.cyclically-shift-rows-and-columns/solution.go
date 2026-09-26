func cyclicShift(n int, grid [][]int, rowShift []int, colShift []int) [][]int {
    for i, k := range rowShift {
        r := make([]int, n)
        for j := range n {
            r[(j-k+n) % n] = grid[i][j]
        }
        grid[i] = r
    }
    for j, k := range colShift {
        c := make([]int, n)
        for i := range n {
            c[(i-k+n) % n] = grid[i][j]
        }
        for i := range n {
            grid[i][j] = c[i]
        }
    }
    return grid
}