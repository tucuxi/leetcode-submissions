func matrixBlockSum(mat [][]int, k int) [][]int {
    answer := make([][]int, len(mat))
    for i := range answer {
        answer[i] = make([]int, len(mat[i]))
        for j := range answer[i] {
            v := 0
            rmax := min(len(mat) - 1, i + k)
            for r := max(0, i - k); r <= rmax; r++ {
                cmax := min(len(mat[i]) - 1, j + k)
                for c := max(0, j - k); c <= cmax; c++ {
                    v += mat[r][c]     
                }
            }
            answer[i][j] = v
        }
    }
    return answer
}
    
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}