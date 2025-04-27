func applyOperations(nums []int) []int {
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i - 1] == nums[i] {
            nums[i - 1] *= 2
            nums[i] = 0
        }
    }
    for p, q := 0, 0; p < n && q < n; {
        if nums[p] != 0 {
            p++
            q = p
        } else {
            if nums[q] != 0 {
                nums[p], nums[q] = nums[q], nums[p]
            }
            q++
        }
    }
    return nums
}