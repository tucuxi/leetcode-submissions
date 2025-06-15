func numberOfAlternatingGroups(colors []int, k int) int {
    res := 0
    run := 1
    previousColor := colors[0]
    for _, color := range append(colors[1:], colors[:k - 1]...) {
        if color != previousColor {
            run++
        } else {
            run = 1
        }
        if run >= k {
            res++
        }
        previousColor = color        
    }
    return res
}