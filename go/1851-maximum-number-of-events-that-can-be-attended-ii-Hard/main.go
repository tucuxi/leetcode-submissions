func maxValue(events [][]int, k int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][0] < events[j][0]
    })
    dp := make([][]int, k + 1)
    for i := range dp {
        dp[i] = make([]int, len(events))
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    
    var dfs func(int, int) int
    dfs = func(index, count int) int {
        if index == len(events) || count == 0 {
            return 0
        }
        if dp[count][index] != -1 {
            return dp[count][index]
        }
        next := index + sort.Search(len(events) - index, func(i int) bool {
            return events[i + index][0] > events[index][1] 
        })
        dp[count][index] = max(dfs(index + 1, count), dfs(next, count - 1) + events[index][2])
        return dp[count][index]
    }
    
    return dfs(0, k) 
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}