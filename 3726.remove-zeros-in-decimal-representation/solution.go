func removeZeros(n int64) int64 {
    res := int64(0)
    factor := int64(1)
    for ; n > 0; n /= 10 {
        if d := n % 10; d > 0 {
            res += d * factor
            factor *= 10
        }
    }
    return res
}