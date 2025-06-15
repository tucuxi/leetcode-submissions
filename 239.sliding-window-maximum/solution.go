func maxSlidingWindow(nums []int, k int) []int {
    var q, res []int
    for i := 0; i < k; i++ {
        for len(q) > 0 && nums[i] >= nums[q[len(q) - 1]] {
            q = q[:len(q) - 1]
        }
        q = append(q, i)
    }
    res = append(res, nums[q[0]])
    for i := k; i < len(nums); i++ {
        if q[0] == i - k {
            q = q[1:]
        }
        for len(q) > 0 && nums[i] >= nums[q[len(q) - 1]] {
            q = q[:len(q) - 1]
        }
        q = append(q, i)
        res = append(res, nums[q[0]])
    }
    return res
}