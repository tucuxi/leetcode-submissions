func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n / 2; i++ {
        k := n - 1 - i
        for j := i; j < k; j++ {
            matrix[i][j], matrix[j][k], matrix[k][n - 1 - j], matrix[n - 1 - j][i] = matrix[n - 1 - j][i], matrix[i][j], matrix[j][k], matrix[k][n - 1 - j]
        }
    }
}