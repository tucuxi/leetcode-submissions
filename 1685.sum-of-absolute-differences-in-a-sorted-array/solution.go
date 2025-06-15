func getSumAbsoluteDifferences(nums []int) []int {
    l, r := 0, 0
    for _, n := range nums {
        r += n
    }

    res := make([]int, len(nums))
    for i := range(nums) {
        r -= nums[i]
        res[i] = r - l + (2*i + 1 - len(nums)) * nums[i]
        l += nums[i]
    }
    return res
}