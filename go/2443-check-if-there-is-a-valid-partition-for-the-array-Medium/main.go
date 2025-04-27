func validPartition(nums []int) bool {
    memo := make([]bool, len(nums))
    visited := make([]bool, len(nums))
    var dp func(int) bool
    dp = func(i int) bool {
        if i == len(nums) {
            return true
        }
        if visited[i] {
            return memo[i]
        }
        if i + 1 < len(nums) && nums[i] == nums[i + 1] {
            memo[i] = dp(i + 2)
        }
        if i + 2 < len(nums) &&
            (nums[i + 1] == nums[i] && nums[i + 2] == nums[i] ||
                 nums[i + 1] == nums[i] + 1 && nums[i + 2] == nums[i] + 2) {
            memo[i] = memo[i] || dp(i + 3)
        }
        visited[i] = true
        return memo[i]
    }
    return dp(0)
}