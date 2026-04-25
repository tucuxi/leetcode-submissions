func isIdealPermutation(nums []int) bool {
    for i, num := range nums {
        if abs(num-i) > 1 {
            return false
        }
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}