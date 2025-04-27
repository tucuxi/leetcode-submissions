func subArrayRanges(nums []int) (res int64) {
    for i := range nums {
        lo, hi := nums[i], nums[i]
        for j := i + 1; j < len(nums); j++ {
            lo = min(lo, nums[j])
            hi = max(hi, nums[j])
            res += int64(hi - lo)
        }
    }
    return
}