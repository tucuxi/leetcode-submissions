func distinctDifferenceArray(nums []int) []int {
    pos := make(map[int]int)
    for _, n := range nums {
        pos[n]++
    }
    pre := make(map[int]int)
    res := make([]int, 0, len(nums))
    for _, n := range nums {
        pre[n]++
        pos[n]--
        if pos[n] == 0 {
            delete(pos, n)
        }
        res = append(res, len(pre) - len(pos))
    }
    return res
}