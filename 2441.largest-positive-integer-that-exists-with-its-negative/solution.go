func findMaxK(nums []int) int {
    h := make(map[int]bool)
    res := -1
    for _, n := range nums {
        if h[-n] {
            res = max(res, n, -n)
        }
        h[n] = true
    }
    return res
}