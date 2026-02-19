func reverseBits(n int) int {
    res := 0
    for range 32 {
        res = (res << 1) | (n & 1)
        n >>= 1
    }
    return res
}