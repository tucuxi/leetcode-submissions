func rob(nums []int) int {
    p, q := 0, 0
    for _, n := range nums {
        if p + n > q {
            p, q = q, p + n
        } else {
            p = q
        }
    }
    return q
}