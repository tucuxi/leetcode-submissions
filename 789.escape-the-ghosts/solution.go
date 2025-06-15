func escapeGhosts(ghosts [][]int, target []int) bool {
    p := abs(target[0]) + abs(target[1])
    for _, g := range ghosts {
        if abs(g[0] - target[0]) + abs(g[1] - target[1]) <= p {
            return false
        }
    }
    return true
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}