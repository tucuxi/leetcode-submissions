func canBeIncreasing(nums []int) bool {
    r := 0
    for i := 1; i < len(nums); i++ {
        if nums[i - 1] >= nums[i] {
            r++
            if i >= 2 && nums[i - 2] >= nums[i] {
                if i < len(nums) - 1 && nums[i - 1] >= nums[i + 1] {
                    return false
                } 
            }
        }
    }
    return r <= 1
}