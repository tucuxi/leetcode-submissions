func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])

    res := make([][]int, m)
    for i := range m {
        res[i] = make([]int, n)
    }

    for r := range m {
        var hl, hr [51]int
        distinctL := 0
        distinctR := 0
        for i, j := r + 1, 1; i < m && j < n; i, j = i + 1, j + 1 {
            v := grid[i][j]
            hr[v]++
            if hr[v] == 1 {
                distinctR++
            }
        }
        res[r][0] = distinctR
        for i, j := r + 1, 1; i < m && j < n; i, j = i + 1, j + 1 {
            vl := grid[i - 1][j - 1]
            hl[vl]++
            if hl[vl] == 1 {
                distinctL++
            }
            v := grid[i][j]
            hr[v]--
            if hr[v] == 0 {
                distinctR--
            }
            res[i][j] = abs(distinctL - distinctR)
        }
    }

    for c := range n {
        var hl, hr [51]int
        distinctL := 0
        distinctR := 0
        for i, j := 1, c + 1; i < m && j < n; i, j = i + 1, j + 1 {
            v := grid[i][j]
            hr[v]++
            if hr[v] == 1 {
                distinctR++
            }
        }
        res[0][c] = distinctR
        for i, j := 1, c + 1; i < m && j < n; i, j = i + 1, j + 1 {
            vl := grid[i - 1][j - 1]
            hl[vl]++
            if hl[vl] == 1 {
                distinctL++
            }
            v := grid[i][j]
            hr[v]--
            if hr[v] == 0 {
                distinctR--
            }
            res[i][j] = abs(distinctL - distinctR)
        }
    }

    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}