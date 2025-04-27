func findingUsersActiveMinutes(logs [][]int, k int) []int {
    h := make(map[int]map[int]bool)
    for _, log := range logs {
        m := h[log[0]]
        if m == nil {
            m = make(map[int]bool)
            h[log[0]] = m
        }
        m[log[1]] = true
    }
    res := make([]int, k)
    for _, m := range h {
        res[len(m) - 1]++
    }
    return res
}