func canMeasureWater(x int, y int, target int) bool {
    v := make(map[[2]int]bool)
    for q := [][2]int{{0, 0}}; len(q) > 0; q = q[1:] {
        if !v[q[0]] { 
            v[q[0]] = true
            j1, j2 := q[0][0], q[0][1]
            if j1 + j2 == target {
                return true
            }
            if j1 > 0 {
                q = append(q, [2]int{0, j2})
            }
            if j2 > 0 {
                q = append(q, [2]int{j1, 0})
            }
            if j1 < x {
                q = append(q, [2]int{x, j2})
            }
            if j2 < y {
                q = append(q, [2]int{j1, y})
            }
            if j1 < x && j2 > 0 {
                p := min(j2, x-j1)
                q = append(q, [2]int{j1+p, j2-p})
            }
            if j1 > 0 && j2 < y {
                p := min(j1, y-j2)
                q = append(q, [2]int{j1-p, j2+p})
            }
        }
    }
    return false
}