func countArrays(original []int, bounds [][]int) int {
    u, v := bounds[0][0], bounds[0][1]
    res := v - u + 1
    for i := 1; i < len(original); i++ {
        d := original[i] - original[i - 1]
        u = max(u + d, bounds[i][0])
        v = min(v + d, bounds[i][1])
        res = min(res, v - u + 1)
    }
    return max(0, res)
}
