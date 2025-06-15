func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        bi := bits(arr[i])
        bj := bits(arr[j])
        if bi == bj {
            return arr[i] < arr[j]
        }
        return bi < bj
    })
    return arr
}

func bits(n int) int {
    c := 0
    for ; n != 0; n &= n - 1 {
        c++
    }
    return c
}