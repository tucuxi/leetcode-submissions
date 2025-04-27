func sortArrayByParity(nums []int) []int {
    for p, q := 0, len(nums) - 1; p < q; {
        switch {
            case nums[p] % 2 == 0: p++
            case nums[q] % 2 == 1: q--
            default: nums[p], nums[q] = nums[q], nums[p]
        }
    }
    return nums
}