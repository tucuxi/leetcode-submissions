func heightChecker(heights []int) int {
    sorted := make([]int, len(heights))
    copy(sorted, heights)
    slices.Sort(sorted)
    res := 0
    for i := range heights {
        if heights[i] != sorted[i] {
            res++
        }
    }
    return res
}