func countMaxOrSubsets(nums []int) int {
    b := 0
    for _, n := range nums {
        b |= n
    }
    return dfs(nums, b)
}

func dfs(rest []int, pre int) int {
    if len(rest) > 0 {
        return dfs(rest[1:], pre) + dfs(rest[1:], pre & ^rest[0])
    }
    if pre == 0 {
        return 1  
    }
    return 0
}