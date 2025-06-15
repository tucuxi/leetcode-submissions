func regionsBySlashes(grid []string) int {
    n := len(grid)
    gridPrime := make([][]bool, 3*n)
    for i := range gridPrime {
        gridPrime[i] = make([]bool, 3*n)
    }
    for i, row := range grid {
        iPrime := 3*i
        for j, cell := range row {
            jPrime := 3*j
            switch cell {
            case '/':
                gridPrime[iPrime][jPrime + 2] = true
                gridPrime[iPrime + 1][jPrime + 1] = true
                gridPrime[iPrime + 2][jPrime] = true
            case '\\':
                gridPrime[iPrime][jPrime] = true
                gridPrime[iPrime + 1][jPrime + 1] = true
                gridPrime[iPrime + 2][jPrime + 2] = true
            }
        }
    }
    regions := 0
    for i := range gridPrime {
        for j := range gridPrime[i] {
            if !gridPrime[i][j] {
                dfs(gridPrime, i, j)
                regions++
            }
        }
    }
    return regions
}

func dfs(g [][]bool, i, j int) {
    if !g[i][j] {
        g[i][j] = true
        if i > 0 {
            dfs(g, i-1, j)
        }
        if i < len(g) - 1 {
            dfs(g, i+1, j)
        }
        if j > 0 {
            dfs(g, i, j-1)
        }
        if j < len(g[i]) - 1 {
            dfs(g, i, j+1)
        }
    }
}