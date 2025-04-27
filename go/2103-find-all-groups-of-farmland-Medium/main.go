func findFarmland(land [][]int) [][]int {
    var res [][]int
    m, n := len(land), len(land[0])
    
    for r1 := 0; r1 < m; r1++ {
        for c1 := 0; c1 < n; {
            if land[r1][c1] < 0 {
                c1 = -land[r1][c1]
            } else if land[r1][c1] == 0 {
                c1++
            } else {
                r2 := r1
                c2 := c1+1
                for ; c2 < n && land[r2][c2] == 1; c2++ {
                }
                for ; r2 < m && land[r2][c1] == 1; r2++ {
                    land[r2][c1] = -c2
                }
                res = append(res, []int{r1, c1, r2-1, c2-1})
                c1 = c2
            }
        } 
    }
    return res
}