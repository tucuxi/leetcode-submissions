func maximumImportance(n int, roads [][]int) int64 {
    degree := make([]int, n)
    for _, r := range roads {
        degree[r[0]]++
        degree[r[1]]++
    }
    sort.Ints(degree)
    
    var res int64
    for i, d := range degree {
        res += int64((i + 1) * d)
    }
    
    return res
}