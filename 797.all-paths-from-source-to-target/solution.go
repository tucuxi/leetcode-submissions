func allPathsSourceTarget(graph [][]int) [][]int {
    res := [][]int{}

    var dfs func([]int, int)
    dfs = func(path []int, node int) {
        path = append(path, node)
        if node == len(graph) - 1 {
            completedPath := make([]int, len(path))
            copy(completedPath, path)
            res = append(res, completedPath)
            return
        }
        for _, neighbor := range graph[node] {
            dfs(path, neighbor)
        }
    }

    dfs([]int{}, 0)
    return res
}