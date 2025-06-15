func matrixReshape(mat [][]int, r int, c int) [][]int {
    if len(mat) * len(mat[0]) != r * c {
        return mat
    }
    res := make([][]int, r)
    k := 0
    for i := range mat {
        for j := range mat[i] {
            if len(res[k]) == c {
                k++
                res[k] = make([]int, 0, c)
            }
            res[k] = append(res[k], mat[i][j])
        }
    }
    return res
}