func isZeroArray(nums []int, queries [][]int) bool {
    changes := make([]int, len(nums) + 1)
    for _, q := range queries {
        changes[q[0]]++
        changes[q[1] + 1]--
    }

    maxValue := 0
    for i := range nums {
        maxValue += changes[i]
        if maxValue < nums[i] {
            return false
        }
    }
    return true
}