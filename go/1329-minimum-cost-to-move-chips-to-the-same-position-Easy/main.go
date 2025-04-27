func minCostToMoveChips(position []int) int {
    odd := 0
    for _, p := range position {
        if p % 2 != 0 {
            odd++
        }
    }
    return min(odd, len(position) - odd)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}