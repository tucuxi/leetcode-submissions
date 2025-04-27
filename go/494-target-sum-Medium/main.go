func findTargetSumWays(nums []int, target int) int {
    memo := make(map[[2]int]int)

    var dfs func(int, int) int
    dfs = func(index, sum int) int {
        if index == len(nums) {
            if sum == target {
                return 1
            }
            return 0
        }
        key := [2]int{index, sum}
        if v, exists := memo[key]; exists {
            return v
        }
        v := dfs(index + 1, sum + nums[index]) + dfs(index + 1, sum - nums[index])
        memo[key] = v
        return v
    }

    return dfs(0, 0)
}