func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    aliceTotal, bobTotal := 0, 0
    bobSet := map[int]bool{}
    for _, a := range aliceSizes {
        aliceTotal += a
    }
    for _, b := range bobSizes {
        bobTotal += b
        bobSet[b] = true
    }
    for _, a := range aliceSizes {
        b := bobTotal - aliceTotal + 2 * a
        if b % 2 == 0 && bobSet[b / 2] {
            return []int{a, b / 2}
        }
    }
    panic("no solution")
}