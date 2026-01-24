func minOperations(nums []int) int {
    h := make(map[int]int)
    for i := len(nums) - 1; i >= 0; i-- {
        n := nums[i]
        h[n]++
        if h[n] == 2 {
            return i / 3 + 1
        }
    }
    return 0
}