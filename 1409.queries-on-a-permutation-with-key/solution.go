func processQueries(queries []int, m int) []int {
    p := make([]int, m)
    for i := 0; i < m; i++ {
        p[i] = i + 1
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        j := 0
        for p[j] != q {
            j++
        }
        res[i] = j
        for ; j > 0; j-- {
            p[j] = p[j - 1]
        }
        p[0] = q
    }
    return res
}