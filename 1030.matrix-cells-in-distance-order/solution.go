type entry struct{
    row int
    col int
    dist int
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
    res := make([][]int, 0)
    
    v := make([][]bool, rows)
    for r := 0; r < rows; r++ {
        v[r] = make([]bool, cols)
    }
    
    q := make([]entry, 0)
    q = append(q, entry{rCenter, cCenter, 0})
    
    for len(q) > 0 {
        p := q[0]
        q = q[1:]
        r, c, d := p.row, p.col, p.dist

        if r < 0 || r >= rows || c < 0 || c >= cols {
            continue
        }
        if v[r][c] {
            continue
        }

        v[r][c] = true
        res = append(res, []int{r, c})

        q = append(q, entry{r-1, c, d+1})
        q = append(q, entry{r+1, c, d+1})
        q = append(q, entry{r, c-1, d+1})
        q = append(q, entry{r, c+1, d+1})
    }
    return res
}