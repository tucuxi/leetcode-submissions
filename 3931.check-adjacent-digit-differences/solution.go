func isAdjacentDiffAtMostTwo(s string) bool {
    for i := 1; i < len(s); i++ {
        if abs(int(s[i]) - int(s[i-1])) > 2 {
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