func transpose(matrix [][]int) [][]int {
    res := make([][]int, len(matrix[0]))
    for i := range res {
        res[i] = make([]int, len(matrix))
        for j := range res[i] {
            res[i][j] = matrix[j][i]
        }
    }
    return res
}