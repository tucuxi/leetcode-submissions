type NumMatrix struct {
    prefix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
    prefix := make([][]int, m)
    for i := 0; i < m; i++ {
        prefix[i] = make([]int, n)
    }
    prefix[0][0] = matrix[0][0]
    for i := 1; i < m; i++ {
        prefix[i][0] = prefix[i - 1][0] + matrix[i][0] 
    }
    for j := 1; j < n; j++ {
        prefix[0][j] = prefix[0][j - 1] + matrix[0][j]
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            prefix[i][j] = prefix[i - 1][j] + prefix[i][j - 1] - prefix[i - 1][j - 1] + matrix[i][j]
        }
    }
    return NumMatrix{prefix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    a := this.prefix[row2][col2]
    if row1 > 0 {
        a -= this.prefix[row1 - 1][col2]
    }
    if col1 > 0 {
        a -= this.prefix[row2][col1 - 1]
    }
    if row1 > 0 && col1 > 0 {
        a += this.prefix[row1 - 1][col1 - 1]
    }
    return a
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */