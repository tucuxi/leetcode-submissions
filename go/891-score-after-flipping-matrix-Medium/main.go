func matrixScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    score := (1 << (n-1)) * m
    for j := 1; j < n; j++ {
        sameBits := 0
        for i := 0; i < m; i++ {
            if grid[i][j] == grid[i][0] {
                sameBits++
            }
        }
        sameBits = max(sameBits, m - sameBits)
        columnScore := (1 << (n-j-1)) * sameBits
        score += columnScore
    }
    return score
}