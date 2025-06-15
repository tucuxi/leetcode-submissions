func maximumUniqueSubarray(nums []int) int {
    res, sum := 0, 0
    i, j := 0, 0
    v := make(map[int]bool)
    for i < len(nums) && j < len(nums) {
        if v[nums[j]] {
            sum -= nums[i]
            v[nums[i]] = false
            i++
        } else {
            sum += nums[j]
            res = max(res, sum)
            v[nums[j]] = true
            j++
        }
    }
    return res
}