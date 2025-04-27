func numOfSubarrays(arr []int) int {
    parity := 0
    odds, evens := 0, 1
    res := 0
    for _, a := range arr {
        parity ^= a & 1
        if parity == 0 {
            res += odds
            evens++
        } else {
            res += evens
            odds++
        }
        res %= 1_000_000_007
    }
    return res
}
