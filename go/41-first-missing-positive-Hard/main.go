func firstMissingPositive(nums []int) int {
    for i := 0; i < len(nums); {
        n := nums[i]
        if n > 0 && n <= len(nums) && n != nums[n - 1] {
            nums[i], nums[n - 1] = nums[n - 1], n
        } else {
            i++
        }
    }
    for i := range nums {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    return len(nums) + 1
}