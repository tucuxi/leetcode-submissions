func minimumRightShifts(nums []int) int {
    i := 1
    for ;; i++ {
        if i == len(nums) {
            return 0
        }
        if nums[i] < nums[i - 1] {
            break
        }
    }
    if nums[i] > nums[0] {
        return -1
    }
    for j := i + 1; ; j++ {
        if j == len(nums) {
            return j - i
        }
        if nums[j] > nums[0] || nums[j] < nums[j - 1] {
            return -1
        }
    }
}