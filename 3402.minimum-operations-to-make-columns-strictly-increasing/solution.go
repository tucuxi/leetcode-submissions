func minimumOperations(grid [][]int) int {
    res := 0
    for c := range grid[0] {
        p := grid[0][c]
        for r := 1; r < len(grid); r++ {
            if grid[r][c] <= p {
                res += p - grid[r][c] + 1
                p++
            } else {
                p = grid[r][c]
            }
        }
    }
    return res
}