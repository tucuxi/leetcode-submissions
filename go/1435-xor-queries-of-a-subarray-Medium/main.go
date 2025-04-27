func xorQueries(arr []int, queries [][]int) []int {
    prefix := make([]int, len(arr) + 1)
    for i, a := range arr {
        prefix[i + 1] = prefix[i] ^ a
    }
    results := make([]int, len(queries))
    for i, q := range queries {
        results[i] = prefix[q[0]] ^ prefix[q[1] + 1]
    }
    return results
}