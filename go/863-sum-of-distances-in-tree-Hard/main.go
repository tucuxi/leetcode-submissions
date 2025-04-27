func sumOfDistancesInTree(n int, edges [][]int) []int {
    graph := make([][]int, n)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    res := make([]int, n)
    count := make([]int, n)

    var postorder func(node, parent int)
    postorder = func(node, parent int) {
        for _, child := range graph[node] {
            if child != parent {
                postorder(child, node)
                res[node] += res[child] + count[child]
                count[node] += count[child]
            }
        }
        count[node]++
    }

    var preorder func(node, parent int)
    preorder = func(node, parent int) {
        for _, child := range graph[node] {
            if child != parent {
                res[child] = res[node] - count[child] + n - count[child]
                preorder(child, node)
            }
        }
    }

    postorder(0, -1)
    preorder(0, -1)
    return res
}