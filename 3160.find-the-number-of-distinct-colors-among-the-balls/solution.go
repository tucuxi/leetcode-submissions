func queryResults(limit int, queries [][]int) []int {
    var res []int

    coloredBalls := make(map[int]int)
    colorCounts := make(map[int]int)
    
    for _, q := range queries {
        b, c := q[0], q[1]
        if currentColor, exists := coloredBalls[b]; exists {
            if colorCounts[currentColor] == 1 {
                delete(colorCounts, currentColor)
            } else {
                colorCounts[currentColor]--
            }
        }
        coloredBalls[b] = c
        colorCounts[c]++
        res = append(res, len(colorCounts))
    }

    return res
}