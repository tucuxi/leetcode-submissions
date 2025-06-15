func search(nums []int, target int) bool {
    left, right := 0, len(nums) - 1
    for left < right && nums[left + 1] == nums[left] {
        left++
    }
    for left < right && nums[right - 1] == nums[right] {
        right--
    }
    for left <= right {
        half := left + (right - left) / 2
        if nums[half] == target {
            return true
        }
        if nums[left] <= nums[half] {
            if target < nums[half] && nums[left] <= target {
                right = half - 1
            } else {
                left = half + 1
            }
        } else {
            if target > nums[half] && nums[right] >= target {
                left = half + 1
            } else {
                right = half - 1
            }
        }
    }
    return false
}