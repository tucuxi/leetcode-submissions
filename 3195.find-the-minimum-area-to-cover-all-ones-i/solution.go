func minimumArea(grid [][]int) int {
    minCol, maxCol := len(grid[0]), 0
    for i := range grid {
        j := 0
        for j < minCol && grid[i][j] == 0 {
            j++
        }
        minCol = j
        j = len(grid[i]) - 1
        for j > maxCol && grid[i][j] == 0 {
            j--
        }
        maxCol = j
    }

    minRow, maxRow := len(grid), 0
    for j := range grid[0] {
        i := 0
        for i < minRow && grid[i][j] == 0 {
            i++
        }
        minRow = i
        i = len(grid) - 1
        for i > maxRow && grid[i][j] == 0 {
            i--
        }
        maxRow = i
    }

    return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}