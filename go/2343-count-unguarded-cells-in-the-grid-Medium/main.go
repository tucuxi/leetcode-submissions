func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    g := make([][]byte, m)
    for i := range m {
        g[i] = make([]byte, n)
    }
    for _, guard := range guards {
        g[guard[0]][guard[1]] = 'G'
    }
    for _, wall := range walls {
        g[wall[0]][wall[1]] = 'W'
    }
    for _, guard := range guards {
        c := guard[1]
        for r := guard[0] - 1; r >= 0 && g[r][c] != 'G' && g[r][c] != 'W'; r-- {
            g[r][c] = 'g'
        }
        for r := guard[0] + 1; r < m && g[r][c] != 'G' && g[r][c] != 'W'; r++ {
            g[r][c] = 'g'
        }
        r := guard[0]
        for c := guard[1] - 1; c >= 0 && g[r][c] != 'G' && g[r][c] != 'W'; c-- {
            g[r][c] = 'g'
        }
        for c := guard[1] + 1; c < n && g[r][c] != 'G' && g[r][c] != 'W'; c++ {
            g[r][c] = 'g'
        }
    }
    unguarded := 0
    for r := range m {
        for c := range n {
            if g[r][c] == 0 {
                unguarded++
            }
        }
        fmt.Println()
    }
    return unguarded
}