func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    needed := make([]int, len(capacity))
    for i := range capacity {
        needed[i] = capacity[i] - rocks[i]
    }
    sort.Ints(needed)
    k := 0
    for k < len(needed) && needed[k] <= additionalRocks {
        additionalRocks -= needed[k]
        k++
    }
    return k
}