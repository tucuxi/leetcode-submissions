func construct2DArray(original []int, m int, n int) [][]int {
    if len(original) != m*n {
        return nil
    }
    res := make([][]int, m)
    for i := range m {
        res[i] = original[i*n : (i+1) * n]
    }
    return res
}