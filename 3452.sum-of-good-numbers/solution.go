func sumOfGoodNumbers(nums []int, k int) int {
    res := 0
    for i, n := range nums {
        if i - k >= 0 && n <= nums[i - k] || i + k < len(nums) && n <= nums[i + k] {
            continue
        }
        res += n
    }
    return res
}