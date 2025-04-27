func minTime(n int, edges [][]int, hasApple []bool) int {
    adj := make([][]int, n)
    for _, edge := range edges {
        adj[edge[0]] = append(adj[edge[0]], edge[1])
        adj[edge[1]] = append(adj[edge[1]], edge[0])
    }

    var dfs func(int, int) int
    dfs = func(vertex int, parent int) int {
        time := 0
        for _, child := range adj[vertex] {
            if child != parent {
                childTime := dfs(child, vertex)
                time += childTime
                if childTime > 0 || hasApple[child] {
                    time += 2
                }
            }
        }
        return time
    }

    return dfs(0, -1)
}