func closestMeetingNode(edges []int, node1 int, node2 int) int {
   walk := func(start int) []int {
        d := make([]int, len(edges))
        for i := range d {
            d[i] = math.MaxInt
        }
        d[start] = 0
        steps := 0
        for p := edges[start]; p != -1 && d[p] == math.MaxInt; p = edges[p] {
            steps++
            d[p] = steps
        }
        return d
    }
    d1, d2 := walk(node1), walk(node2)
    minValue, minIndex := math.MaxInt, -1
    for i := range d1 {
        if d1[i] == math.MaxInt || d2[i] == math.MaxInt {
            continue
        }
        if candidate := max(d1[i], d2[i]); candidate < minValue {
            minValue = candidate
            minIndex = i
        }
    }
    return minIndex
}