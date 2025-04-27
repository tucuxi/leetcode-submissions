func findSubsequences(nums []int) [][]int {
    res := make([][]int, 0)

    var dfs func(int, []int)
    dfs = func(start int, path []int) {
        if len(path) > 1 {
            solution := make([]int, len(path))
            copy(solution, path)
            res = append(res, solution)
        }
        used := map[int]bool{}
        for i := start; i < len(nums); i++ {
            if len(path) > 0 && nums[i] < path[len(path) - 1] || used[nums[i]] {
                continue
            }
            used[nums[i]] = true
            dfs(i + 1, append(path, nums[i]))
        }
    }

    dfs(0, make([]int, 0, len(nums)))
    return res
}