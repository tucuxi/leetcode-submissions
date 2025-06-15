func smallestNumber(n int) int {
    res := n
    for m := 1; m < n; m <<= 1 {
        res |= m 
    }
    return res
}