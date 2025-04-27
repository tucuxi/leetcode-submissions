func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
    vl := make([]struct{value, label int}, len(values))
    for i := range vl {
        vl[i].value = values[i]
        vl[i].label = labels[i]
    }
    sort.Slice(vl, func(i, j int) bool { return vl[i].value < vl[j].value })
    score := 0
    used := make(map[int]int)
    for i := len(vl) - 1; numWanted > 0 && i >= 0; i-- {
        if used[vl[i].label] < useLimit {
            used[vl[i].label]++
            score += vl[i].value
            numWanted--
        }
    }
    return score
}