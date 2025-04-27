func putMarbles(weights []int, k int) int64 {
    return maxMinDiff(pairWeights(weights), k)
}

func pairWeights(weights []int) []int {
    pw := make([]int, len(weights) - 1)
    for i := range pw {
        pw[i] = weights[i] + weights[i + 1]
    }
    sort.Ints(pw)
    return pw
}

func maxMinDiff(pw []int, k int) int64 {
    res := int64(0)
    for i := 0; i < k - 1; i++ {
        res += int64(pw[len(pw) - 1 - i] - pw[i])
    }
    return res
}