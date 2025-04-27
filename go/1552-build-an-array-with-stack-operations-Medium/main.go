func buildArray(target []int, n int) []string {
    var res []string
    k := 0
    for i := 1; k < len(target) && i <= n; i++ {
        res = append(res, "Push")
        if i == target[k] {
            k++
        } else {
            res = append(res, "Pop")
        }       
    }
    return res
}