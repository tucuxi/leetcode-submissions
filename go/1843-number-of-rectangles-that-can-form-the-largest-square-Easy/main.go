func countGoodRectangles(rectangles [][]int) int {
    maxLen := 0
    count := 0
    for _, rect := range rectangles {
        var sq int
        if rect[0] < rect[1] {
            sq = rect[0]
        } else {
            sq = rect[1]
        }
        if sq > maxLen {
            maxLen = sq
            count = 0
        }
        if sq == maxLen {
            count++
        }
    }
    return count
}