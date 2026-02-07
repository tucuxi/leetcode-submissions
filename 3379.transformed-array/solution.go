func constructTransformedArray(nums []int) []int {
    res := make([]int, len(nums))
    n := len(nums)
    for i, num := range nums {
        res[i] = nums[(i + num % n + n) % n]
    }
    return res
}