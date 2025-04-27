func shuffle(nums []int, n int) []int {
    res := make([]int, 0, 2 * n)
    for i := 0; i < n; i++ {
        res = append(res, nums[i], nums[i + n])
    }
    return res
}