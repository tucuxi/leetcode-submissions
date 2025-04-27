func maximumTripletValue(nums []int) int64 {
    var (
        res int64
        maxI, maxDiff int
    )
    for _, num := range nums {
        res = max(res, int64(num) * int64(maxDiff))
        maxI = max(maxI, num)
        maxDiff = max(maxDiff, maxI - num)
    }
    return res
}