func kWeakestRows(mat [][]int, k int) []int {
    h := make([]int, len(mat))
    x := make([]int, len(mat))
    for i := range mat {
        n := 0
        for n < len(mat[i]) && mat[i][n] == 1 {
            n++
        }
        h[i] = n
        x[i] = i
    }
    sort.Slice(x, func(i, j int) bool {
        xi, xj := x[i], x[j]
        return h[xi] < h[xj] || h[xi] == h[xj] && xi < xj
    })
    return x[:k]
}
