func minimumAbsDifference(arr []int) [][]int {
    slices.Sort(arr)

    m := math.MaxInt
    for i := 1; i < len(arr); i++ {
        m = min(m, arr[i] - arr[i - 1])
    }

    var res [][]int
    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i - 1] == m {
            res = append(res, []int{arr[i - 1], arr[i]})
        }
    }

    return res
}