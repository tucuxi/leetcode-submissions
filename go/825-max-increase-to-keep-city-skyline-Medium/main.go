func maxIncreaseKeepingSkyline(grid [][]int) int {
    h := make([]int, len(grid))
    v := make([]int, len(grid))
    for i := range grid {
        for j := range grid[i] {
            b := grid[i][j]
            if b > h[i] {
                h[i] = b
            }
            if b > v[j] {
                v[j] = b
            }
        }
    }
    res := 0
    for i := range grid {
        for j := range grid[i] {
            res += min(h[i], v[j]) - grid[i][j]
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}