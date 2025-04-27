func twoCitySchedCost(costs [][]int) int {
    slices.SortFunc(costs, func(a, b []int) int {
        da := a[0] - a[1]
        if da < 0 {
            da = -da
        }
        db := b[0] - b[1]
        if db < 0 {
            db = -db
        }
        return db - da
    })
    a, b := 0, 0
    total := 0
    n := len(costs) / 2
    for _, c := range costs {
        if b == n || a < n && c[0] < c[1] {
            a++
            total += c[0]
        } else {
            b++
            total += c[1]
        }
    }
    return total
}