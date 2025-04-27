func containsCycle(grid [][]byte) bool {
    v := make([][]bool, len(grid))
    for i := range grid {
        v[i] = make([]bool, len(grid[i]))
    }
    
    var cycle func(int, int, int, int, byte) bool
    cycle = func(rPrevious, cPrevious, rCurrent, cCurrent int, value byte) bool {
        v[rCurrent][cCurrent] = true
        for _, step := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
            rNext, cNext := rCurrent + step[0], cCurrent + step[1]
            if rNext >= 0 && rNext < len(grid) && cNext >= 0 && cNext < len(grid[rNext]) {
                if rNext == rPrevious && cNext == cPrevious || grid[rNext][cNext] != value {
                    continue
                }
                if v[rNext][cNext] || cycle(rCurrent, cCurrent, rNext, cNext, value) {
                    return true
                }
            }
        }
        return false
    }
    
    for i := range grid {
        for j := range grid[i] {
            if !v[i][j] {
                if cycle(-1, -1, i, j, grid[i][j]) {
                    return true
                }
            }
        }
    }
    return false
}