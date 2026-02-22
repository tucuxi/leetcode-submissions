func leastBricks(wall [][]int) int {
    gaps := make(map[int]int)

    for _, w := range wall {
        p := 0
        for i := range len(w) - 1 {
            p += w[i]
            gaps[p]++
        }
    }

    maxCount := 0

    for _, count := range gaps {
        maxCount = max(maxCount, count) 
    }

    return len(wall) - maxCount
}