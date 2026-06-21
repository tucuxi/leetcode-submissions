func consecutiveSetBits(n int) bool {
    count := 0
    for ; n > 0; n >>= 1 {
        if n&3 == 3 {
            count++
        }
    }
    return count==1
}