func diagonalSort(mat [][]int) [][]int {
    pos := func(startRow int, startCol int, i int) (int, int, bool) {
        row := startRow + i
        col := startCol + i
        ok := row < len(mat) && col < len(mat[0])
        return row, col, ok
    }
    for row := range mat {
        var diag []int
        for i := 0;; i++ {
            r, c, ok := pos(row, 0, i)
            if !ok {
                break
            }
            diag = append(diag, mat[r][c])
        }
        sort.Ints(diag)
        for i := 0;; i++ {
            r, c, ok := pos(row, 0, i)
            if !ok {
                break
            }
            mat[r][c] = diag[i]
        }
    }
    for col := range mat[0] {
        var diag []int
        for i := 0;; i++ {
            r, c, ok := pos(0, col, i)
            if !ok {
                break
            }
            diag = append(diag, mat[r][c])
        }
        sort.Ints(diag)
        for i := 0;; i++ {
            r, c, ok := pos(0, col, i)
            if !ok {
                break
            }
            mat[r][c] = diag[i]
        }        
    }
    return mat    
}