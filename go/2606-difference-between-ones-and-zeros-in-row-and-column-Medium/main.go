func onesMinusZeros(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    onesRow := make([]int, m)
    onesCol := make([]int, n)
    for i := range grid {
        for j := range grid[i] {
            onesRow[i] += grid[i][j]
            onesCol[j] += grid[i][j]
        }
    }
    diff := make([][]int, m)
    for i := range diff {
        diff[i] = make([]int, n)
        for j := range diff[i] {
            diff[i][j] = 2 * (onesRow[i] + onesCol[j]) - (m + n)
        }
    }
    return diff
}