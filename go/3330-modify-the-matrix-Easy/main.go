func modifiedMatrix(matrix [][]int) [][]int {
    mc := make([]int, len(matrix[0]))  
    for j := range mc {
        mc[j] = matrix[0][j]
        for i := 1; i < len(matrix); i++ {
            mc[j] = max(mc[j], matrix[i][j])
        }
    }
    for i := range matrix {
        for j := range mc {
            if matrix[i][j] == -1 {
                matrix[i][j] = mc[j]
            }
        }
    }
    return matrix
}