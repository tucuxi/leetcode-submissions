func finishTime(n int, edges [][]int, baseTime []int) int64 {
    children := make([][]int, n)
    for _, e := range edges {
        u := e[0]
        v := e[1]
        children[u] = append(children[u], v)
    }

    var dfs func(int) int64
    dfs = func(u int) int64 {
        if children[u] == nil {
            return int64(baseTime[u])
        }

        var (
            earliest int64 = math.MaxInt64
            latest   int64 = math.MinInt64
        )
        
        for _, v := range children[u] {
            finishTime := dfs(v)
            earliest = min(earliest, finishTime)
            latest = max(latest, finishTime)
        }

        ownDuration := latest - earliest + int64(baseTime[u])
        return latest + ownDuration
    }

    return dfs(0)
}