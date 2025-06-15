func numSubmatrixSumTarget(matrix [][]int, target int) int {
    m := len(matrix)
    n := len(matrix[0])
    p := make([][]int, m+1)
    
    for i := 0; i <= m; i++ {
        p[i] = make([]int, n+1)
    }

    for i := 0; i < m; i++ {
		s := 0
		for j := 0; j < n; j++ {
			s += matrix[i][j]
			p[i+1][j+1] = s + p[i][j+1]
		}
	}
    
    res := 0
    for i1 := 1; i1 <= m; i1++ {
        for i2 := 0; i2 < i1; i2++ {
            for j1 := 1; j1 <= n; j1++ {
                for j2 := 0; j2 < j1; j2++ {
                    s := p[i1][j1] - p[i1][j2] - p[i2][j1] + p[i2][j2]
                    if s == target {
                        res++
                    }
                }
            }
        }
    }
    
    return res
}