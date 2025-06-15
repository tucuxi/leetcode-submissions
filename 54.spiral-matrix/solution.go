func spiralOrder(matrix [][]int) []int {
    w, h := len(matrix[0]), len(matrix)
    r, c := 0, -1
    dr, dc := 1, 1
    res := make([]int, 0, w * h)
    for w > 0 && h > 0 {
        for i := 0; i < w; i++ {
            c += dc
            res = append(res, matrix[r][c])
        }
        dc = -dc
        h--
        for j := 0; j < h; j++ {
            r += dr
            res = append(res, matrix[r][c])
        }
        dr = -dr
        w--
    }
    return res
}