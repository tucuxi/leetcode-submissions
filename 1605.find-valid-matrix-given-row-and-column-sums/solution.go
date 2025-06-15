func restoreMatrix(rowSum []int, colSum []int) [][]int {
    m := make([][]int, len(rowSum))
    for r := range m {
        m[r] = make([]int, len(colSum))
        for c := range m[r] {
            a := min(rowSum[r], colSum[c])
            rowSum[r] -= a
            colSum[c] -= a
            m[r][c] = a
        }
    }
    return m
}