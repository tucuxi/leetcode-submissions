func generate(numRows int) [][]int {
    tri := make([][]int, numRows)
    for i := range tri {
        tri[i] = make([]int, i + 1)
        tri[i][0] = 1
        tri[i][i] = 1
        for j := 1; j < i; j++ {
            tri[i][j] = tri[i - 1][j] + tri[i - 1][j - 1]
        }
    }
    return tri
}