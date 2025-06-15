func projectionArea(grid [][]int) int {
    xy := 0
    for y := range grid {
        for x := range grid[y] {
            if grid[y][x] > 0 {
                xy++
            }
        }
    }
    xz := 0
    for x := range grid[0] {
        m := 0
        for y := range grid {
            if grid[y][x] > m {
                m = grid[y][x]
            }
        }
        xz += m
    }
    yz := 0
    for y := range grid {
        m := 0
        for x := range grid[y] {
            if grid[y][x] > m {
                m = grid[y][x]
            }
        }
        yz += m
    }
    return xy + xz + yz
}