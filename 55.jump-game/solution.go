func canJump(nums []int) bool {
    maxDistance := 0
    for i := range nums {
        if maxDistance < i {
            return false
        }
        if maxJump := i + nums[i]; maxJump > maxDistance {
            maxDistance = maxJump
        }
        if maxDistance >= len(nums) - 1 {
            return true
        }
    }
    return false
}