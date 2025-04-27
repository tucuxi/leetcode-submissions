func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    var res [][]int
    dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    r := rStart
    c := cStart
    d := 0
    l := 2
    k := 0
    for len(res) < rows*cols {
        if r >= 0 && r < rows && c >= 0 && c < cols {
            res = append(res, []int{r, c})
        }
        if k == l/2 {
            k = 0
            l++
            d = (d+1) % len(dir)
        }
        k++
        r += dir[d][0]
        c += dir[d][1]
    }
    return res
}