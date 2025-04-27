func circularGameLosers(n int, k int) []int {
    v := make([]bool, n)
    f := 0
    i := 1
    for !v[f] {
        v[f] = true
        f = (f + i * k) % n
        i++
    }
    res := make([]int, 0)
    for i := range v {
        if !v[i] {
            res = append(res, i + 1)
        }
    }
    return res
}