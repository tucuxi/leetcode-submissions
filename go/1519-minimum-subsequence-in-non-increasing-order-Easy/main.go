func minSubsequence(nums []int) []int {
    res := []int{}
    p, q := 0, 0
    for _, n := range nums {
        p += n
    }
    sort.Ints(nums)
    for k := len(nums) - 1; k >= 0 && q <= p; k-- {
        p -= nums[k]
        q += nums[k]
        res = append(res, nums[k])
    }
    return res
}