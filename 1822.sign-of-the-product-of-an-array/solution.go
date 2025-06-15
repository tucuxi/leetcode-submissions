func arraySign(nums []int) int {
    res := 1
    for _, n := range nums {
        if n < 0 {
            res = -res
        } else if n == 0 {
            return 0
        }
    }
    return res
}