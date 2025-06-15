func maximumDetonation(bombs [][]int) int {
    g := make([][]int, len(bombs))
    r := func(i, j int) bool {
        dx := bombs[i][0] - bombs[j][0]
        dy := bombs[i][1] - bombs[j][1]
        return dx * dx + dy * dy <= bombs[i][2] * bombs[i][2]
    }
    for i := range bombs {
        for j := range bombs {
            if i != j && r(i, j) {
                g[i] = append(g[i], j)
            }
        }
    }
    res := 0
    for i := range bombs {
        res = max(res, bfs(g, i))
    }
    return res
}

func bfs(g [][]int, i int) int {
    q := []int{i}
    visited := make([]bool, len(g))
    visited[i] = true
    n := 1
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        for _, v := range g[u] {
            if !visited[v] {
                visited[v] = true
                n++
                q = append(q, v)
            }
        }
    }
    return n
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}