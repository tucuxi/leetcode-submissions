func canMakeSquare(grid [][]byte) bool {
    for i := 0; i < len(grid) - 1; i++ {
        for j := 0; j < len(grid[i]) - 1; j++ {
            b := 0
            if grid[i][j] == 'B' {
                b++
            }
            if grid[i][j+1] == 'B' {
                b++
            }
            if grid[i+1][j] == 'B' {
                b++
            }
            if grid[i+1][j+1] == 'B' {
                b++
            }
            if b != 2 {
                return true
            } 
        }
    }
    return false
}