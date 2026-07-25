func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    group := make([]int, n)
    g := 0

    for i := 1; i < n; i++ {
        if nums[i] - nums[i-1] > maxDiff {
            g++
        }
        group[i] = g
    }

    res := make([]bool, len(queries))
    
    for i, q := range queries {
        res[i] = group[q[0]] == group[q[1]]
    }
    return res
}