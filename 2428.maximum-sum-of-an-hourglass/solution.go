func maxSum(grid [][]int) int {
    max := 0
    for r := 1; r <= len(grid) - 2; r++ {
        for c := 1; c <= len(grid[r]) - 2; c++ {
            s := grid[r - 1][c - 1] + grid[r - 1][c] + grid[r - 1][c + 1] + grid[r][c] + grid[r + 1][c - 1] + grid[r + 1][c] + grid[r + 1][c + 1]
            if s > max {
                max = s
            }
        }
    }
    return max
}