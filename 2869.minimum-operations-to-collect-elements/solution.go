func minOperations(nums []int, k int) int {
    s := make([]bool, k + 1)
    w := 0
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] <= k && !s[nums[i]] {
            s[nums[i]] = true
            w++
            if w == k {
                return len(nums) - i
            }
        }
    }
    panic("no solution")
}