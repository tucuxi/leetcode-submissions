func replaceElements(arr []int) []int {
    max := -1
    for i := len(arr) - 1; i >= 0; i-- {
        cur := arr[i]
        arr[i] = max
        if cur > max {
            max = cur
        }
    }
    return arr
}