func shortestPathLength(graph [][]int) int {
    var (
        queue [][]int
        visited = make([]map[int]bool, len(graph))
        all = (1 << len(graph)) - 1
    )
    for i := range graph {
        mask := 1 << i
        queue = append(queue, []int{i, mask})
        visited[i] = map[int]bool{mask: true}
    }
    for l := 0; ; l++ {
        for _, e := range queue {
            queue = queue[1:]
            if e[1] == all {
                return l
            }
            for _, v := range graph[e[0]] {
                mask := e[1] | (1 << v)
                if !visited[v][mask] {
                    visited[v][mask] = true
                    queue = append(queue, []int{v, mask})
                }
            } 
        }
    }
}