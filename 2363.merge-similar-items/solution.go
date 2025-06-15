func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    weights := map[int]int{}
    for i := range items1 {
        weights[items1[i][0]] += items1[i][1]
    }
    for i := range items2 {
        weights[items2[i][0]] += items2[i][1]
    }
    ret := [][]int{}
    for v, w := range weights {
        ret = append(ret, []int{v, w})
    }
    sort.Slice(ret, func(i, j int) bool { return ret[i][0] < ret[j][0] })
    return ret
}