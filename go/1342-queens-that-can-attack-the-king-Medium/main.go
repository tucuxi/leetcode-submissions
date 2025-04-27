func queensAttacktheKing(queens [][]int, king []int) [][]int {
    var b [8][8]bool
    var res [][]int
    
    for _, q := range queens {
        b[q[0]][q[1]] = true
    }
    
    for _, q := range queens {
        x, y := q[0], q[1]
        dx, dy := sgn(king[0] - x), sgn(king[1] - y)
        for {
            x += dx
            y += dy
            if x < 0 || x >= 8 || y < 0 || y >= 8 || b[x][y] {
                break
            }
            if x == king[0] && y == king[1] {
                res = append(res, []int{q[0], q[1]})
                break
            }
        }
    }
    
    return res
    
}

func sgn(a int) int {
    if a < 0 {
        return -1
    }
    if a > 0 {
        return 1
    }
    return 0
}