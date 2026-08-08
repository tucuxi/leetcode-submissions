func aggregateTimeSeries(series1 [][]int, series2 [][]int) [][]int {
    res := make([][]int, 0, len(series1) + len(series2))
    i, j := 0, 0
    for i < len(series1) && j < len(series2) {
        if series1[i][0] < series2[j][0] {
            res = append(res, []int{series1[i][0], series1[i][1] + series2[j][1]})
            i++
        } else if series1[i][0] > series2[j][0] {
            res = append(res, []int{series2[j][0], series1[i][1] + series2[j][1]})
            j++
        } else {
            res = append(res, []int{series1[i][0], series1[i][1] + series2[j][1]})
            i++
            j++
        }
    }
    if i < len(series1) {
        res = append(res, series1[i:]...)
    }
    if j < len(series2) {
        res = append(res, series2[j:]...)
    }
    return res
}