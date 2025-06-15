func minimumOperations(nums []int) int {
    u := map[int]bool{}
    for _, n := range nums {
        if n > 0 {
            u[n] = true
        }
    }
    return len(u)
}