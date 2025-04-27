func zigzagTraversal(grid [][]int) []int {
    var res []int
    n := len(grid[0])

    for r := range grid {
        if r % 2 == 0 {
            for c := 0; c < n; c += 2 {
                res = append(res, grid[r][c])
            }
        } else {
            for c := n - 1 - n % 2; c >= 0; c -= 2 {
                res = append(res, grid[r][c])
            }
        }
    }
    return res
}
