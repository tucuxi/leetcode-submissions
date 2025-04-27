func checkStraightLine(coordinates [][]int) bool {
    x0 := coordinates[0][0]
    y0 := coordinates[0][1]
    dx1 := coordinates[1][0] - x0
    dy1 := coordinates[1][1] - y0
    for i := 2; i < len(coordinates); i++ {
        dxi := coordinates[i][0] - x0
        dyi := coordinates[i][1] - y0
        if dx1 * dyi != dy1 * dxi {
            return false  
        } 
    }
    return true
}