func containsDuplicate(nums []int) bool {
    h := map[int]bool{}
    for _, n := range nums {
        if h[n] {
            return true
        }
        h[n] = true
    }
    return false
}