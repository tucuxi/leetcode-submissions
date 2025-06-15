func countValidSelections(nums []int) int {
    res := 0
    total := 0
    for _, n := range nums {
        total += n
    }
    prefixSum := 0
    for _, n := range nums {
        prefixSum += n
        if n == 0 {
            d := total - 2 * prefixSum
            if d == 0 {
                res += 2
            } else if d == 1 || d == -1 {
                res++
            }
        }
    }
    return res
}