func largestPathValue(colors string, edges [][]int) int {
    adj, inDegree := makeGraph(len(colors), edges)
    nodes := []int{}
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
        dp[u][colors[u] - 'a']++
        res = max(res, dp[u][colors[u] - 'a'])
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

func makeGraph(n int, edges [][]int) (map[int][]int, []int) {
    adj := make(map[int][]int)
    in := make([]int, n)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        in[e[1]]++
    }
    return adj, in    
}

func max(a ...int) int {
    res := a[0]
    for _, n := range a {
        if n > res {
            res = n
        }
    }
    return res
}