func minBitFlips(start int, goal int) int {
    c := 0
    for x := start ^ goal; x != 0; x &= x-1 {
        c++
    }
    return c
}