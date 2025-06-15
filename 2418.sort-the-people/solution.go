func sortPeople(names []string, heights []int) []string {
    idx := make([]int, len(names))
    for i := range idx {
        idx[i] = i
    }
    slices.SortFunc(idx, func(i, j int) int {
        return heights[j] - heights[i]
    })
    res := make([]string, len(names))
    for i := range idx {
        res[i] = names[idx[i]]
    }
    return res
}