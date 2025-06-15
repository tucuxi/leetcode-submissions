func firstCompleteIndex(arr []int, mat [][]int) int {
    m, n := len(mat), len(mat[0])
    r := make([]int, m)
    c := make([]int, n)
    h := make([][]int, m * n + 1)
    for i := range mat {
        for j := range mat[i] {
            h[mat[i][j]] = []int{i, j}
        }
    }
    for k := range arr {
        a := arr[k]
        i := h[a][0]
        r[i]++
        if r[i] == n {
            return k
        }
        j := h[a][1]
        c[j]++
        if c[j] == m {
            return k
        }
    }
    panic("no solution")
}