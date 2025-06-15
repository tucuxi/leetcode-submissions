func resultsArray(nums []int, k int) []int {
    res := make([]int, len(nums) - k + 1)
    p := nums[0]
    r := 0
    for i, n := range nums {
        if n != p + 1 {
            r = 0
        }
        r++
        if j := i - k + 1; r >= k {
            res[j] = n
        } else if j >= 0 {
            res[j] = -1
        }
        p = n
    }
    return res
}