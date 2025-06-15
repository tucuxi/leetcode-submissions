func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    row := make([]bool, m)
    col := make([]bool, n)
    for i := range m {
        for j := range n {
            if matrix[i][j] == 0 {
                row[i] = true
                col[j] = true
            }
        }
    }
    for i := range m {
        for j := range n {
            if row[i] || col[j] {
                matrix[i][j] = 0
            }
        }
    }
}