func findClosest(x int, y int, z int) int {
    switch d := abs(x - z) - abs(y - z); {
        case d < 0: return 1
        case d > 0: return 2
        default: return 0
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}