func findMinHeightTrees(n int, edges [][]int) []int {
    inDegree := make([]int, n)
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        inDegree[u]++
        inDegree[v]++
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    var q []int
    vis := make([]bool, n)
    for u := range inDegree {
        if inDegree[u] <= 1 {
            q = append(q, u)
            vis[u] = true
        }
    }
    for {
        k := len(q)
        for i := 0; i < k; i++ {
            u := q[i]
            for _, v := range adj[u] {
                if !vis[v] {
                    inDegree[v]--
                    if inDegree[v] <= 1 {
                        q = append(q, v)
                        vis[v] = true
                    }
                }
            }
        }
        if len(q) == k {
            return q
        }
        q = q[k:]
    }
}