func countServers(grid [][]int) int {
    rowCount := make([]int, len(grid))
    colCount := make([]int, len(grid[0]))
    
    for i := range grid {
        for j, c := range grid[i] {
            rowCount[i] += c
            colCount[j] += c
        }
    }

    res := 0
    for i := range grid {
        for j, c := range grid[i] {
            if c == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
                res++
            }
        }
    }
    return res
}