func countAsterisks(s string) int {
    res := 0
    inpair := false
    for _, c := range s {
        switch c {
        case '|':
            inpair = !inpair
        case '*':
            if !inpair {
                res++
            }
        }
    }
    return res
}