func groupThePeople(groupSizes []int) [][]int {
    h := make(map[int][]int)
    res := make([][]int, 0)
    for i, s := range groupSizes {
        h[s] = append(h[s], i)
        if len(h[s]) == s {
            res = append(res, h[s])
            delete(h, s)
        }
    }
    return res
}