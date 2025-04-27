func isPathCrossing(path string) bool {
    var l [2]int
    v := make(map[[2]int]bool)
    v[l] = true
    for _, p := range path {
        switch p {
        case 'N': l[0]--
        case 'S': l[0]++
        case 'W': l[1]--
        case 'E': l[1]++
        }
        if v[l] {
            return true
        }
        v[l] = true
    }
    return false
}