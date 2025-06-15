func successfulPairs(spells []int, potions []int, success int64) []int {
    res := make([]int, len(spells))
    sort.Ints(potions)
    for i := range spells {
        j := sort.Search(len(potions), func(j int) bool {
            s := int64(spells[i]) * int64(potions[j])
            return s >= success
        })
        res[i] = len(potions) - j
    }
    return res
}