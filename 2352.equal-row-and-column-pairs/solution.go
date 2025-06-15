func equalPairs(grid [][]int) int {
    count := 0
    for i := range grid {
        columns:
        for j := range grid {
            for k := range grid {
                if grid[i][k] != grid[k][j] {
                    continue columns
                }
            }
            count++
        }
    }
    return count
}