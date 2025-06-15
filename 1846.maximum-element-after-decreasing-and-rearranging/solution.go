func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    res := 0
    sort.Ints(arr)
    for _, a := range arr {
        if a > res {
            res++
        }
    }
    return res
}