func numMagicSquaresInside(grid [][]int) int {
    res := 0
    for i := 0; i < len(grid) - 2; i++ {
  
        outer:
        for j := 0; j < len(grid[0]) - 2; j++ {
            var h [10]bool
            var rowSums, colSums [3]int 

            for k := 0; k < 3; k++ {
                for l := 0; l < 3; l++ {
                    v := grid[i+k][j+l]
                    if v < 1 || v > 9 || h[v] {
                        continue outer
                    }
                    h[v] = true
                    rowSums[k] += v
                    colSums[l] += v
                }
            }
            if !allShouldBe(15, rowSums) ||
                !allShouldBe(15, colSums) ||
                grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2] != 15 ||
                grid[i+2][j] + grid[i+1][j+1] + grid[i][j+2] != 15 {
                continue outer
            }
            res++
        }
    }
    return res
}

func allShouldBe(expected int, values [3]int) bool {
    for _, v := range values {
        if v != expected {
            return false
        }
    }
    return true
}