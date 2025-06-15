func minOperations(nums []int, k int) int {
    set := make(map[int]bool)

    for _, n := range nums {
        if n < k {
            return -1
        }
        set[n] = true
    }

    if set[k] {
        return len(set) - 1
    }
    return len(set)
}