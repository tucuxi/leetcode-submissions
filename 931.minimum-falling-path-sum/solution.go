func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    p, q := make([]int, n), make([]int, n)
    for j := 0; j < n; j++ {
        p[j] = matrix[n - 1][j]
    }
    for i := n - 2; i >= 0; i-- {
        for j := 0; j < n; j++ {
            min := p[j]
            if j > 0 && p[j - 1] < min {
                min = p[j - 1]
            }
            if j + 1 < n && p[j + 1] < min {
                min = p[j + 1]
            }
            q[j] = matrix[i][j] + min 
        }
        p, q = q, p
    }
    min := math.MaxInt
    for j := 0; j < n; j++ {
        if p[j] < min {
            min = p[j]
        }
    }
    return min
}