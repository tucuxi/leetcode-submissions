func maxEqualRowsAfterFlips(matrix [][]int) int {
    h := make(map[string]int)
    res := 0

    for _, row := range matrix {
        var b strings.Builder

        for _, cell := range row {
            b.WriteByte('0' + byte(cell ^ row[0]))
        }
        s := b.String()
        h[s]++
        res = max(res, h[s])
    }
    return res
}
