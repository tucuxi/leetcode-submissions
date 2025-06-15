func kthSmallest(matrix [][]int, k int) int {
    l := len(matrix) - 1
    return matrix[0][0] + sort.Search(matrix[l][l] - matrix[0][0] + 1, func(x int) bool {
        c := 0
        for i := range matrix {
            p := sort.Search(len(matrix), func(j int) bool {
                return matrix[i][j] - matrix[0][0] > x
            })
            c += p
        }
        return c >= k
    })
}