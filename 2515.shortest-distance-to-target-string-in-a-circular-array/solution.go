func closestTarget(words []string, target string, startIndex int) int {
    n := len(words)
    for i := range n/2 + 1 {
        if words[(startIndex-i+n) % n] == target || words[(startIndex+i) % n] == target {
            return i
        }
    }
    return -1
}