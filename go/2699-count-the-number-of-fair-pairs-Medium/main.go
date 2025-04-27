func countFairPairs(nums []int, lower int, upper int) int64 {
    slices.Sort(nums)
    return lowerBound(nums, upper + 1) - lowerBound(nums, lower)
}

func lowerBound(nums []int, bound int) int64 {
    var res int64
    i := 0
    j := len(nums) - 1
    for i < j {
        if nums[i] + nums[j] < bound {
            res += int64(j - i)
            i++
        } else {
            j--
        }
    }
    return res
}