func highestPeak(isWater [][]int) [][]int {
    var q [][2]int
    cells := make([][]int, len(isWater))
    for i := range cells {
        cells[i] = make([]int, len(isWater[i]))
        for j := range cells[i] {
            if isWater[i][j] == 0 {
                cells[i][j] = -1
            } else {
                q = append(q, [2]int{i, j})
            }
        }
    }
    d := 0
    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            y, x := q[0][0], q[0][1]
            q = q[1:]
            if cells[y][x] > 0 || cells[y][x] == 0 && d > 0 {
                continue
            }
            cells[y][x] = d
            if y > 0 && cells[y-1][x] == -1 {
                q = append(q, [2]int{y-1, x})
            }
            if y+1 < len(cells) && cells[y+1][x] == -1 {
                q = append(q, [2]int{y+1, x})
            }
            if x > 0 && cells[y][x-1] == -1 {
                q = append(q, [2]int{y, x-1})
            }
            if x+1 < len(cells[y]) && cells[y][x+1] == -1 {
                q = append(q, [2]int{y, x+1})
            }
        }
        d++
    }
    return cells
}