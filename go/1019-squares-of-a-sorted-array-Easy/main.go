func sortedSquares(nums []int) []int {
    var res []int

    r, _ := slices.BinarySearch(nums, 0)
    l := r-1

    for l >= 0 || r < len(nums) {
        if l < 0 {
            res = append(res, nums[r] * nums[r])
            r++
        } else if r == len(nums) {
            res = append(res, nums[l] * nums[l])
            l--
        } else if nums[r] < -nums[l] {
            res = append(res, nums[r] * nums[r])
            r++
        } else {
            res = append(res, nums[l] * nums[l])
            l--
        }
    }
    return res
}