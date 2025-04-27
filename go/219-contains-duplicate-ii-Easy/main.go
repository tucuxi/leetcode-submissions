func containsNearbyDuplicate(nums []int, k int) bool {
    ind := map[int]int{}
    for i, n := range nums {
        j, ok := ind[n]
        if ok && i - j <= k {
            return true
        }
        ind[n] = i
    }
    return false
}
