func minimumReplacement(nums []int) int64 {
    var res int64
    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] > nums[i + 1] {
            k := int64(nums[i] + nums[i + 1] - 1) / int64(nums[i + 1])
            res += k - 1
            nums[i] /= int(k)
        }
    }
    return res
}