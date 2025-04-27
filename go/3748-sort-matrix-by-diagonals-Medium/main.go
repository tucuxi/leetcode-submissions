func sortMatrix(grid [][]int) [][]int {
    n := len(grid)

    for r := range n {
        sort.Sort(sort.Reverse(diagonal{grid, r, 0}))
    }
    for c := 1; c < n; c++ {
        sort.Sort(diagonal{grid, 0, c})
    }

    return grid
}

type diagonal struct {
    grid     [][]int
    startRow int
    startCol int
}

func (d diagonal) Len() int {
    return len(d.grid) - max(d.startRow, d.startCol)
}

func (d diagonal) Less(i, j int) bool {
    return d.value(i) < d.value(j) 
}

func (d diagonal) Swap(i, j int) {
    ri, ci := d.indices(i)
    rj, cj := d.indices(j)
    d.grid[ri][ci], d.grid[rj][cj] = d.grid[rj][cj], d.grid[ri][ci]
}

func (d diagonal) indices(i int) (int, int) {
    return d.startRow + i, d.startCol + i
}

func (d diagonal) value(i int) int {
    return d.grid[d.startRow + i][d.startCol + i]
}