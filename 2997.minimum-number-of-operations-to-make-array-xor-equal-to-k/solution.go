func minOperations(nums []int, k int) int {
    x := k
    for _, n := range nums {
        x ^= n
    }
    c := 0
    for ; x != 0; x &= x-1 {
        c++
    }
    return c
}