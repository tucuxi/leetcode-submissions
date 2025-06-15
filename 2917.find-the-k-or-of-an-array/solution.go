func findKOr(nums []int, k int) int {
    res := 0
    for i := 0; i < 32; i++ {
        b := 1 << i
        c := 0
        for _, n := range nums {
            if n & b != 0 {
                c++
                if c == k {
                    res |= b
                    break
                }
            }
        }
    }
    return res
}