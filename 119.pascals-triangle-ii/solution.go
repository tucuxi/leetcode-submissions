func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    a := make([]int, rowIndex + 1)
    a[0] = 1
    a[rowIndex] = 1
    b := getRow(rowIndex - 1)
    for i := 1; i < rowIndex; i++ {
        a[i] = b[i - 1] + b[i]
    }
    return a
}