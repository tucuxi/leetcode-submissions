func diagonalSum(mat [][]int) int {
    sum := 0
    for i, j := 0, len(mat) - 1; i < len(mat); i, j = i + 1, j - 1 {
        sum += mat[i][i]
        if i != j {
            sum += mat[i][j]
        }
    }
    return sum
}