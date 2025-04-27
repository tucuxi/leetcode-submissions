func hasTrailingZeros(nums []int) bool {
    c := 0
    for _, n := range nums {
        if n & 1 == 0 {
            c++
        }
    }
    return c >= 2
}
