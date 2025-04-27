func removeDuplicates(nums []int) int {
    last := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[last] {
            last++
            nums[last] = nums[i]
        }
    }
    return last + 1
}