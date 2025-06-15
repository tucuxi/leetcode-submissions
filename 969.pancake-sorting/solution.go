func pancakeSort(arr []int) []int {
    var res []int

    for n := len(arr); n > 1; n-- {
        k := 0
        for i := range n {
            if arr[i] > arr[k] {
                k = i
            }
        }
        if k > 0 {
            res = append(res, k + 1)
            for i := range k / 2 + 1 {
                arr[i], arr[k - i] = arr[k - i], arr[i]
            }
        }
        res = append(res, n)
        for i := range n / 2 {
            arr[i], arr[n - 1 - i] = arr[n - 1 - i], arr[i]
        }
    }
    return res
}

