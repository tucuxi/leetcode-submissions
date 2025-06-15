func minImpossibleOR(nums []int) int {
    m := 0
    for _, n := range nums {
        if n & (n - 1) == 0 {
            m |= n
        }
    }
    res := 1
    for m & 1 != 0 {
        res <<= 1
        m >>= 1
    }
    return res
}