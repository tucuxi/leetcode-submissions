func arrayRankTransform(arr []int) []int {
    index := make([]int, len(arr))
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(a, b int) int { return arr[a] - arr[b] })
    rank := 0
    previous := math.MinInt
    for i := range arr {
        if arr[index[i]] != previous {
            rank++
            previous = arr[index[i]]
        }
        arr[index[i]] = rank
    }
    return arr
}