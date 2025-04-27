func singleNumber(nums []int) []int {
    a, b := 0, 0
    for _, n := range nums {
        a ^= n
    }
    x := a & -a
    for _, n := range nums {
        if n & x != 0 {
            b ^= n
        }
    }
    return []int{b, a ^ b}
}