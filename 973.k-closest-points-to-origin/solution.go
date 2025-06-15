func kClosest(points [][]int, k int) [][]int {
    sort.Slice(points, func(i, j int) bool {
        return sq(points[i][0]) + sq(points[i][1]) < sq(points[j][0]) + sq(points[j][1])
    })
    return points[:k]
}

func sq(a int) int {
    return a * a
}