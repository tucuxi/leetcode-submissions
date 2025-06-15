func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    res := make([][]int, n-2)
    for i := 0; i < n-2; i++ {
        res[i] = make([]int, n-2)
        for j := 0; j < n-2; j++ {
            m := grid[i][j]
            for k := i+2; k >= i; k-- {
                for l := j+2; l >= j; l-- {
                    if grid[k][l] > m {
                        m = grid[k][l]
                    }
                }
            }
            res[i][j] = m
        }
    }
    return res
}