func numSpecial(mat [][]int) int {
    rows := make([]int, len(mat))
    cols := make([]int, len(mat[0]))
    for i := range mat {
        for j := range mat[0] {
            rows[i] += mat[i][j]
            cols[j] += mat[i][j]
        }
    }
    res := 0
    for i := range rows {
        if rows[i] != 1 {
            continue
        }
        for j := range cols {
            if mat[i][j] == 1 && cols[j] == 1 {
                res++
            }
        }
    }
    return res
}