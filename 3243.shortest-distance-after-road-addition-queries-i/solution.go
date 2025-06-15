func shortestDistanceAfterQueries(n int, queries [][]int) []int {
    succ := make([][]int, n)
    for i := range n - 1 {
        succ[i] = []int{i + 1}
    }
    
    res := make([]int, 0, len(queries))
    for _, query := range queries {
        a, b := query[0], query[1]
        succ[a] = append(succ[a], b)

        steps := 0
        v := make(map[int]bool)
        v[0] = true

        bfs:
        for queue := []int{0}; len(queue) > 0; steps++ {
            k := len(queue)
            for i := range k {
                node := queue[i]
                if node == n - 1 {
                    break bfs
                }
                for _, s := range succ[node] {
                    if !v[s] {
                        v[s] = true
                        queue = append(queue, s)
                    }
                }
            }
            queue = queue[k:]
        }
        res = append(res, steps)
    }
    return res
}