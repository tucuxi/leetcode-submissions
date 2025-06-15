func minimumOperations(nums []int) int {
    operations := 0
    for _, n := range nums {
        rest := n % 3
        operations += min(rest, 3 - rest)
    }
    return operations
}