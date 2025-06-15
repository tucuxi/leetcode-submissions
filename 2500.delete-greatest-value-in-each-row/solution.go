func deleteGreatestValue(grid [][]int) int {
    res := 0
    for i := range grid {
        sort.Ints(grid[i])
    }
    for j := len(grid[0]) - 1; j >= 0; j-- {
        max := 0
        for i := range grid {
            if grid[i][j] > max {
                max = grid[i][j]
            }
        }
        res += max
    }
    return res
}