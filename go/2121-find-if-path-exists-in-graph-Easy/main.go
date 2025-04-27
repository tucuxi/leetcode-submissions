func validPath(n int, edges [][]int, source int, destination int) bool {
    neighbors := make([][]int, n)
    for _, edge := range edges {
        neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
        neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
    }
    visited := make([]bool, n)

    var dfs func(int) bool
    dfs = func(node int) bool {
        visited[node] = true
        if node == destination {
            return true
        }
        for _, neighbor := range neighbors[node] {
            if !visited[neighbor] && dfs(neighbor) {
                return true
            }
        }
        return false
    }

    return dfs(source)
}