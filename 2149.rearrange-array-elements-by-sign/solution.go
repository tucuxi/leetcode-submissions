func rearrangeArray(nums []int) []int {
    res := make([]int, len(nums))
    p, q := 0, 1
    for _, n := range nums {
        if n > 0 {
            res[p] = n
            p += 2
        } else {
            res[q] = n
            q += 2
        }
    }
    return res
}