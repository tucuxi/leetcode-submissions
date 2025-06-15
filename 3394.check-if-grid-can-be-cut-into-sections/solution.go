func checkValidCuts(n int, rectangles [][]int) bool {

    check := func(dim int) bool {
        slices.SortFunc(rectangles, func(a, b []int) int {
            return a[dim] - b[dim]
        })
        gaps := 0
        maxEnd := rectangles[0][dim + 2]
        for _, r := range rectangles[1:] {
            if maxEnd <= r[dim] {
                gaps++
            }
            maxEnd = max(maxEnd, r[dim + 2])
        }
        return gaps >= 2
    }

    return check(0) || check(1)
}
