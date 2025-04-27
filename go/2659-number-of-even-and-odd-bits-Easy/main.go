func evenOddBit(n int) []int {
    res := make([]int, 2)
    i := 0
    for n != 0 {
        if n & 1 == 1 {
            res[i]++
        }
        i ^= 1
        n >>= 1
    }
    return res
}