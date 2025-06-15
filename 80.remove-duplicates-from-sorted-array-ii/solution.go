func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
        return len(nums)
    }
    k := 2
    for i := 2; i < len(nums); i++ {
        if nums[i] > nums[k - 2] || nums[i] > nums[k - 1] {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}