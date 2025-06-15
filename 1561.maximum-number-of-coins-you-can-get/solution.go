func maxCoins(piles []int) int {
    sort.Ints(piles)
    res := 0
    b := 0
    i := len(piles) - 2
    for b < i {
        res += piles[i]
        b++
        i -= 2
    }
    return res
}