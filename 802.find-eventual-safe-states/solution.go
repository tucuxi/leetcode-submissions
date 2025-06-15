func eventualSafeNodes(graph [][]int) []int {
    from := make([][]int, len(graph))
    outDegree := make([]int, len(graph))
    res := make([]int, 0)
    q := make([]int, 0)
    for u, adj := range graph {
        outDegree[u] = len(adj)
        if len(adj) == 0 {
            q = append(q, u)
        }
        for _, v := range adj {
            from[v] = append(from[v], u)
        }
    }
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        res = append(res, u)
        for _, v := range from[u] {
            outDegree[v]--
            if outDegree[v] == 0 {
                q = append(q, v)
            }
        }
    }   
    sort.Ints(res)
    return res
}