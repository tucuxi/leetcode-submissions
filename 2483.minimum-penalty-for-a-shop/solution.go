func bestClosingTime(customers string) int {
    curP := 0
    minP := 0
    minT := 0
    for t, c := range customers {
        if c == 'Y' {
            curP--
        } else {
            curP++
        }
        if curP < minP {
            minP = curP
            minT = t + 1
        }
    }
    return minT
}