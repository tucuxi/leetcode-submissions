func findOriginalArray(changed []int) []int {
    if len(changed) % 2 == 1 {
        return []int{}
    }
    sort.Ints(changed)
    h := map[int]int{}
    for _, v := range changed {
        h[v]++
    }
    res := []int{}
    for i, k := 0, 0; i < len(changed); i++ {
        n := changed[i]
        if h[n] > 0 {
            h[n]--
            if h[2*n] == 0 {
                return []int{}
            }
            h[2*n]--
            res = append(res, n)
            k++
        }
    }
    if len(res) != len(changed) / 2 {
        return []int{}
    }
    return res
}