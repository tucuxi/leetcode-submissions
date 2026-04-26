func maxDistance(colors []int) int {
    res := 0
    cl, cr := colors[0], colors[len(colors) - 1]
    for i := range colors {
        if colors[i] != cl {
            res = max(res, i)
        }
        if colors[i] != cr {
            res = max(res, len(colors) -1 - i)
        }
    }
    return res
}