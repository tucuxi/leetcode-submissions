func garbageCollection(garbage []string, travel []int) int {
    res := 0
    var dist [3]int
    for i, g := range garbage {
        res += len(g)
        for j, b := range check(g) {
            if b {
                res += dist[j]
                dist[j] = 0
            }
            if i < len(travel) {
                dist[j] += travel[i]
            }
        }
    }
    return res
}

func check(s string) [3]bool {
    var res [3]bool
    for _, r := range s {
        switch r {
        case 'M': res[0] = true
        case 'P': res[1] = true
        case 'G': res[2] = true
        }
    }
    return res
}