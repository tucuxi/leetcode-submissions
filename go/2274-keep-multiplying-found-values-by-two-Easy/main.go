func findFinalValue(nums []int, original int) int {
    numSet := map[int]bool{}
    for _, n := range nums {
        numSet[n] = true
    }
    for numSet[original] {
        original *= 2
    } 
    return original
}