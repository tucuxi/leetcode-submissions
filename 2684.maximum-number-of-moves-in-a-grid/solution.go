func maxMoves(grid [][]int) int {
    n := len(grid[0])
    q := make(map[int]bool)
    moves := 0

    for i := range grid {
        q[i] = true
    }
    for j := 1; j < n; j++ {
        p := make(map[int]bool)

        for r := range q {
            for i := max(0, r - 1); i <= min(len(grid) - 1, r + 1); i++ {
                if grid[i][j] > grid[r][j - 1] {
                    p[i] = true
                }
            }
        }
        if len(p) == 0 {
            break
        }
        q = p
        moves++
    }
    return moves
}