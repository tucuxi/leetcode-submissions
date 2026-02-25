func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        bi := bits.OnesCount(uint(arr[i]))
        bj := bits.OnesCount(uint(arr[j]))
        if bi == bj {
            return arr[i] < arr[j]
        }
        return bi < bj
    })
    return arr
}