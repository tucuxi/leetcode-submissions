func subarraySum(nums []int) int {
    pre := make([]int, len(nums) + 1)
    for i := range nums {
        pre[i + 1] = pre[i] + nums[i]
    }

    res := 0
    for i := range nums {
        res += pre[i + 1] - pre[max(0, i - nums[i])]
    }

    return res
}