func canBeEqual(target []int, arr []int) bool {
    h := make(map[int]int)
    for _, n := range target {
        h[n]++
    }
    for _, n := range arr {
        if h[n] == 1 {
            delete(h, n)
        } else  {
            h[n]--
        }
    }
    return len(h) == 0
}