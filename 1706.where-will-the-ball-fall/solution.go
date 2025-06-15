func findBall(grid [][]int) []int {
    m, n := len(grid), len(grid[0])
    answer := make([]int, n)

    outer:
    for i := range answer {
        answer[i] = -1
        col := i
        for row := range m {
            if grid[row][col] == 1 {
                col++
                if col == n || grid[row][col] == -1 {
                    continue outer
                }
            } else {
                col--
                if col < 0 || grid[row][col] == 1 {
                    continue outer
                }
            }
        }
        answer[i] = col
    }

    return answer
}