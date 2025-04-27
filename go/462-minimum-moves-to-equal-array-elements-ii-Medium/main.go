func minMoves2(nums []int) int {
    sort.Ints(nums)
    k := len(nums) / 2
    m := nums[k]
    res := 0
    for i := 0; i < k; i++ {
        res += m - nums[i]
    }
    for i := k; i < len(nums); i++ {
        res += nums[i] - m
    }
    return res
}