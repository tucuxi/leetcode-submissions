func findRotation(mat [][]int, target [][]int) bool {
    n := len(mat)
    excl := [4]bool{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            excl[0] = excl[0] || mat[i][j] != target[i][j]
            excl[1] = excl[1] || mat[i][j] != target[j][n - i - 1]
            excl[2] = excl[2] || mat[i][j] != target[n - i - 1][n - j - 1]
            excl[3] = excl[3] || mat[i][j] != target[n - j - 1][i]
        }
    }
    return !excl[0] || !excl[1] || !excl[2] || !excl[3]
}