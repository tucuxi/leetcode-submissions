func maximumDifference(nums []int) int {
    res, min := -1, nums[0]
    for _, n := range nums {
        if min < n {
            res = max(res, n - min)
        } else {
            min = n
        }
    }
    return res
}