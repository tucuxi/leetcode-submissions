func findColumnWidth(grid [][]int) []int {
    res := make([]int, len(grid[0]))
    for i := range grid {
        for j := range grid[i] {
            if k := strconv.Itoa(grid[i][j]); len(k) > res[j] {
                res[j] = len(k)
            }
        }
    }
    return res
}