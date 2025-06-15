func minFlips(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    rowFlips := 0
    for i := range m {
        for j := range n/2 {
            if grid[i][j] != grid[i][n-1-j] {
                rowFlips++
            }
        }
    }
    colFlips := 0
    for j := range n {
        for i := range m/2 {
            if grid[i][j] != grid[m-1-i][j] {
                colFlips++
            }
        }
    }
    return min(rowFlips, colFlips)
}