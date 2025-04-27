func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    intersection := make([][]int, 0, len(firstList))
    i, j := 0, 0
    for i < len(firstList) && j < len(secondList) {
        a, b := firstList[i][0], firstList[i][1]
        c, d := secondList[j][0], secondList[j][1]
        switch {
        case b < c:
            i++
        case a > d:
            j++
        case b > d:
            intersection = append(intersection, []int{max(a, c), d})
            j++
        default:
            intersection = append(intersection, []int{max(a, c), b})
            i++
        }

    }
    return intersection
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}