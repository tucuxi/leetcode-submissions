func minGroups(intervals [][]int) int {
    changes := make([][2]int, 0, 2 * len(intervals))
    for _, interval := range intervals {
        changes = append(changes, [2]int{interval[0], 1}, [2]int{interval[1] + 1, -1})
    }
    slices.SortFunc(changes, func(a, b [2]int) int {
        if d := a[0] - b[0]; d != 0 {
            return d
        }
        return a[1] - b[1]
    })
    currentSum := 0
    maxSum := 0
    for _, change := range changes {
        currentSum += change[1]
        maxSum = max(currentSum, maxSum)
    }
    return maxSum
}