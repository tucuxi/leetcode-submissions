func isNStraightHand(hand []int, groupSize int) bool {
    slices.Sort(hand)
    h := make(map[int]int)
    for _, card := range hand {
        h[card]++
    }
    for _, card := range hand {
        if h[card] == 0 {
            continue
        }
        for i := card; i < card + groupSize; i++ {
            if h[i] == 0 {
                return false
            }
            h[i]--
        }
    }
    return true
}
