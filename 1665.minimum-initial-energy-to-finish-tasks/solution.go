func minimumEffort(tasks [][]int) int {
    slices.SortFunc(tasks, func(a, b []int) int {
        return (a[1] - a[0]) - (b[1] - b[0])
    })
    res := 0
    for _, t := range tasks {
        if res + t[0] > t[1] {
            res += t[0]
        } else {
            res = t[1]
        }
    }
    return res
}