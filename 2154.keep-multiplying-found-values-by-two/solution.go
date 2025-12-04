func findFinalValue(nums []int, original int) int {
    v := make([]bool, 1001)
    for _, n := range nums {
        v[n] = true
    }
    n := original
    for n < len(v) && v[n] {
        n += n
    }
    return n
}