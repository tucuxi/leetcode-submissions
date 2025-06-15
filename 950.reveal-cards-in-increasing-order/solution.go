func deckRevealedIncreasing(deck []int) []int {
    index := makeIndex(len(deck))
    res := make([]int, len(deck))
    sort.Ints(deck)
    for _, c := range deck {
        res[index[0]] = c
        index = index[1:]
        if len(index) > 0 {
            index = append(index[1:], index[0])
        }
    }
    return res
}

func makeIndex(n int) (index []int) {
    for i := 0; i < n; i++ {
        index = append(index, i)
    }
    return
}