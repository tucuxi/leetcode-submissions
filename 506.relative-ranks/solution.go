var medals = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
    index := make([]int, len(score))
    for i := range index {
        index[i] = i
    }
    slices.SortFunc(index, func(i, j int) int { return score[j] - score[i] })
    res := make([]string, len(score))
    for i := 0; i < min(len(index), 3); i++ {
        res[index[i]] = medals[i]
    }
    for i := 3; i < len(index); i++ {
        res[index[i]] = strconv.Itoa(i+1)
    }
    return res
}