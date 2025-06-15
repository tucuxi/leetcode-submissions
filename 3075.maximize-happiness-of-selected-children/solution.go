func maximumHappinessSum(happiness []int, k int) int64 {
    var res int64
    
    slices.Sort(happiness)
    for i := 0; i < k; i++ {
        h := happiness[len(happiness) - 1 - i] - i
        if h <= 0 {
            break
        }
        res += int64(h)
    }

    return res
}