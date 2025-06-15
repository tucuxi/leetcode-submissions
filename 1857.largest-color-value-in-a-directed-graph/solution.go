func largestPathValue(colors string, edges [][]int) int {
    adj, inDegree := func() (map[int][]int, []int) {
        adj := make(map[int][]int)
        in := make([]int, len(colors))
        for _, e := range edges {
            adj[e[0]] = append(adj[e[0]], e[1])
            in[e[1]]++
        }
        return adj, in    
    }()

    var nodes []int

    for i, d := range inDegree {
        if d == 0 {
            nodes = append(nodes, i)
        }
    }
    res := 0
    dp := make([][26]int, len(colors))
    for len(nodes) > 0 {
        u := nodes[0]
        nodes = nodes[1:]
        col := colors[u] - 'a'
        dp[u][col]++
        res = max(res, dp[u][col])
        for _, v := range adj[u] {
            for c := range dp[u] {
                dp[v][c] = max(dp[v][c], dp[u][c])
            }
            inDegree[v]--
            if inDegree[v] == 0 {
                nodes = append(nodes, v)
            }
        }
    }
    for _, d := range inDegree {
        if d > 0 {
            return -1
        }
    }
    return res
}