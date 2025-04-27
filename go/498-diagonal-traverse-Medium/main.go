func findDiagonalOrder(mat [][]int) []int {
    res := make([]int, 0, len(mat) * len(mat[0]))
    dr, dc := -1, 1
    r, c := 0, 0
    for len(res) < cap(res) {
        res = append(res, mat[r][c])
        r += dr
        c += dc
        if r < 0 && c >= len(mat[0]) {
            dr, dc = 1, -1
            r, c = 1, len(mat[0]) - 1
        } else if r >= len(mat) && c < 0 {
            dr, dc = -1, 1
            r, c = len(mat) - 1, 1
        } else if r < 0 {
            dr, dc = 1, -1
            r = 0
        } else if r >= len(mat) {
            dr, dc = -1, 1
            r, c = len(mat) - 1, c + 2
        } else if c < 0 {
            dr, dc = -1, 1
            c = 0
        } else if c >= len(mat[0]) {
            dr, dc = 1, -1
            r, c = r + 2, len(mat[0]) - 1
        }
    }
    return res
}