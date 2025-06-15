func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])

    total := func() int {
        sum := 0
        for _, row := range grid {
            for _, cell := range row {
                sum += cell
            }
        }
        return sum
    }()

    checkVertical := func() bool {
        prefixVertical := 0
        for i := range m - 1 {
            for _, cell := range grid[i] {
                prefixVertical += cell
            }
            if 2*prefixVertical == total {
                return true
            }
        }
        return false
    }

    checkHorizontal := func() bool {
        prefixHorizontal := 0
        for j := range n - 1 {
            for i := range m {
                prefixHorizontal += grid[i][j]
            }
            if 2*prefixHorizontal == total {
                return true
            }
        }
        return false
    }

    return checkVertical() || checkHorizontal()
}