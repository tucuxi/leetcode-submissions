func areSimilar(mat [][]int, k int) bool {
    n := len(mat[0])
    for r := range mat {
        var o int
        if r % 2 == 1 {
            o = k % n
        } else {
            o = n - k % n
        }
        for c := range mat[r] {
            if mat[r][c] != mat[r][(c + o) % n] {
                return false
            }
        }
    }
    return true
}